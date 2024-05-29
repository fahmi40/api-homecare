package medicalofficer

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
	app.DB().RegisterTable("main", MedicalOfficer{})
	app.DB().MigrateTable(tx, "main", app.Setting{})
	tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&MedicalOfficer{})

	app.Server().AddMiddleware(app.Test().NewCtx([]string{
		"medical_officers.detail",
		"medical_officers.list",
		"medical_officers.create",
		"medical_officers.edit",
		"medical_officers.delete",
	}))
	app.Server().AddRoute("/medical_officers", "POST", REST().Create, nil)
	app.Server().AddRoute("/medical_officers", "GET", REST().Get, nil)
	app.Server().AddRoute("/medical_officers/:id", "GET", REST().GetByID, nil)
	app.Server().AddRoute("/medical_officers/:id", "PUT", REST().UpdateByID, nil)
	app.Server().AddRoute("/medical_officers/:id", "PATCH", REST().PartiallyUpdateByID, nil)
	app.Server().AddRoute("/medical_officers/:id", "DELETE", REST().DeleteByID, nil)
}

// getTestMedicalOfficerID returns an available MedicalOfficer ID.
func getTestMedicalOfficerID() string {
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
		description:  "Get empty list of MedicalOfficer",
		method:       "GET",
		path:         "/medical_officers",
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"count":0,"results":[]}`,
	},
	{
		description:  "Create MedicalOfficer with minimum payload",
		method:       "POST",
		path:         "/medical_officers",
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"name":"Kilogram"}`,
		expectedCode: http.StatusCreated,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Get MedicalOfficer by ID",
		method:       "GET",
		path:         "/medical_officers/" + getTestMedicalOfficerID(),
		token:        app.TestFullAccessToken,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilogram"}`,
	},
	{
		description:  "Update MedicalOfficer by ID",
		method:       "PUT",
		path:         "/medical_officers/" + getTestMedicalOfficerID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Update MedicalOfficer by ID","name":"KG"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"KG"}`,
	},
	{
		description:  "Partially update MedicalOfficer by ID",
		method:       "PATCH",
		path:         "/medical_officers/" + getTestMedicalOfficerID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Partially Update MedicalOfficer by ID","name":"Kilo Gram"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"name":"Kilo Gram"}`,
	},
	{
		description:  "Delete MedicalOfficer by ID",
		method:       "DELETE",
		path:         "/medical_officers/" + getTestMedicalOfficerID(),
		token:        app.TestFullAccessToken,
		bodyRequest:  `{"reason":"Delete MedicalOfficer by ID"}`,
		expectedCode: http.StatusOK,
		expectedBody: `{"code":200}`,
	},
}

// TestMedicalOfficerREST tests the REST API of MedicalOfficer data with specified scenario.
func TestMedicalOfficerREST(t *testing.T) {
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

// BenchmarkMedicalOfficerREST tests the REST API of MedicalOfficer data with specified scenario.
func BenchmarkMedicalOfficerREST(b *testing.B) {
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
