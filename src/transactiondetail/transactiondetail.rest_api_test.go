package transactiondetail

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"gorm.io/gorm"

	"home-care/app"
)

// prepareTest prepares the test.
func prepareTest(tb testing.TB) {
	app.Test()
	tx := app.Test().Tx
	app.DB().RegisterTable("main", TransactionDetail{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&TransactionDetail{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"transaction_details.detail",
		"transaction_details.list",
		"transaction_details.create",
		"transaction_details.edit",
		"transaction_details.delete",
	}))
	app.Server().AddRoute("/transaction_details", "POST", REST().Create, nil)
	app.Server().AddRoute("/transaction_details", "GET", REST().Get, nil)
	app.Server().AddRoute("/transaction_details/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/transaction_details/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/transaction_details/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/transaction_details/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestTransactionDetailID returns an available TransactionDetail ID.
func getTestTransactionDetailID() string {
	return "todo"
}

// tests is test scenario.
var tests = []struct {
	description  string // description of the test case
	method       string // method to test
	path         string // route path to test
	token        string // token to test
	bodyRequest  string // body to test
	expectedCode int    // expected HTTP status code
	expectedBody string // expected body response
}{
	{
		description:  "Get empty list of TransactionDetail",
		method:       "GET",
		path:         "/transaction_details",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create TransactionDetail with minimum payload",
		method:       "POST",
		path:         "/transaction_details",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get TransactionDetail by ID",
		method:       "GET",
		path:         "/transaction_details/" + getTestTransactionDetailID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update TransactionDetail by ID",
		method:       "PUT",
		path:         "/transaction_details/" + getTestTransactionDetailID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update TransactionDetail by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update TransactionDetail by ID",
		method:       "PATCH",
		path:         "/transaction_details/" + getTestTransactionDetailID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update TransactionDetail by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete TransactionDetail by ID",
		method:       "DELETE",
		path:         "/transaction_details/" + getTestTransactionDetailID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete TransactionDetail by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestTransactionDetailREST tests the REST API of TransactionDetail data with specified scenario.
func TestTransactionDetailREST(t *testing.T) {
	prepareTest(t)

	// Iterate through test single test cases
	for _, test := range tests {

		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
		req.Header.Add("Authorization", "Bearer "+test.token)
		req.Header.Add("Content-Type", "application/json")

		// Perform the request plain with the app, the second argument is a request latency (set to -1 for no latency)
		res, err := app.Server().Test(req)

		// Verify if the status code is as expected
		utils.AssertEqual(t, nil, err, "app.Server().Test(req)")
		utils.AssertEqual(t, test.expectedCode, res.StatusCode, test.description)

		// Verify if the body response is as expected
		body, err := io.ReadAll(res.Body)
		utils.AssertEqual(t, nil, err, "io.ReadAll(res.Body)")
		app.Test().AssertMatchJSONElement(t, []byte(test.expectedBody), body, test.description)
		res.Body.Close()
	}
}

// BenchmarkTransactionDetailREST tests the REST API of TransactionDetail data with specified scenario.
func BenchmarkTransactionDetailREST(b *testing.B) {
	b.ReportAllocs()
	prepareTest(b)
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			req := httptest.NewRequest(test.method, test.path, strings.NewReader(test.bodyRequest))
			req.Header.Add("Authorization", "Bearer "+test.token)
			req.Header.Add("Content-Type", "application/json")
			app.Server().Test(req)
		}
	}
}
