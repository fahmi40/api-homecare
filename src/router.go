package src

import (
	"home-care/app"
	"home-care/src/medicalofficer"
	"home-care/src/medicaltreatment"
	"home-care/src/registration"
	"home-care/src/schedule"
	"home-care/src/transaction"
	"home-care/src/transactiondetail"
	"home-care/src/user"
	// import : DONT REMOVE THIS COMMENT
)

func Router() *routerUtil {
	if router == nil {
		router = &routerUtil{}
		router.Configure()
		router.isConfigured = true
	}
	return router
}

var router *routerUtil

type routerUtil struct {
	isConfigured bool
}

func (r *routerUtil) Configure() {
	app.Server().AddRoute("/api/version", "GET", app.VersionHandler, nil)

	app.Server().AddRoute("/api/users", "POST", user.REST().Create, user.OpenAPI().Create())
	app.Server().AddRoute("/api/users", "GET", user.REST().Get, user.OpenAPI().Get())
	app.Server().AddRoute("/api/users/{id}", "GET", user.REST().GetByID, user.OpenAPI().GetByID())
	app.Server().AddRoute("/api/users/{id}", "PUT", user.REST().UpdateByID, user.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/users/{id}", "PATCH", user.REST().PartiallyUpdateByID, user.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/users/{id}", "DELETE", user.REST().DeleteByID, user.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/medical_officers", "POST", medicalofficer.REST().Create, medicalofficer.OpenAPI().Create())
	app.Server().AddRoute("/api/medical_officers", "GET", medicalofficer.REST().Get, medicalofficer.OpenAPI().Get())
	app.Server().AddRoute("/api/medical_officers/{id}", "GET", medicalofficer.REST().GetByID, medicalofficer.OpenAPI().GetByID())
	app.Server().AddRoute("/api/medical_officers/{id}", "PUT", medicalofficer.REST().UpdateByID, medicalofficer.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/medical_officers/{id}", "PATCH", medicalofficer.REST().PartiallyUpdateByID, medicalofficer.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/medical_officers/{id}", "DELETE", medicalofficer.REST().DeleteByID, medicalofficer.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/medical_treatments", "POST", medicaltreatment.REST().Create, medicaltreatment.OpenAPI().Create())
	app.Server().AddRoute("/api/medical_treatments", "GET", medicaltreatment.REST().Get, medicaltreatment.OpenAPI().Get())
	app.Server().AddRoute("/api/medical_treatments/{id}", "GET", medicaltreatment.REST().GetByID, medicaltreatment.OpenAPI().GetByID())
	app.Server().AddRoute("/api/medical_treatments/{id}", "PUT", medicaltreatment.REST().UpdateByID, medicaltreatment.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/medical_treatments/{id}", "PATCH", medicaltreatment.REST().PartiallyUpdateByID, medicaltreatment.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/medical_treatments/{id}", "DELETE", medicaltreatment.REST().DeleteByID, medicaltreatment.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/registrations", "POST", registration.REST().Create, registration.OpenAPI().Create())
	app.Server().AddRoute("/api/registrations", "GET", registration.REST().Get, registration.OpenAPI().Get())
	app.Server().AddRoute("/api/registrations/{id}", "GET", registration.REST().GetByID, registration.OpenAPI().GetByID())
	app.Server().AddRoute("/api/registrations/{id}", "PUT", registration.REST().UpdateByID, registration.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/registrations/{id}", "PATCH", registration.REST().PartiallyUpdateByID, registration.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/registrations/{id}", "DELETE", registration.REST().DeleteByID, registration.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/schedules", "POST", schedule.REST().Create, schedule.OpenAPI().Create())
	app.Server().AddRoute("/api/schedules", "GET", schedule.REST().Get, schedule.OpenAPI().Get())
	app.Server().AddRoute("/api/schedules/{id}", "GET", schedule.REST().GetByID, schedule.OpenAPI().GetByID())
	app.Server().AddRoute("/api/schedules/{id}", "PUT", schedule.REST().UpdateByID, schedule.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/schedules/{id}", "PATCH", schedule.REST().PartiallyUpdateByID, schedule.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/schedules/{id}", "DELETE", schedule.REST().DeleteByID, schedule.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/transactions", "POST", transaction.REST().Create, transaction.OpenAPI().Create())
	app.Server().AddRoute("/api/transactions", "GET", transaction.REST().Get, transaction.OpenAPI().Get())
	app.Server().AddRoute("/api/transactions/{id}", "GET", transaction.REST().GetByID, transaction.OpenAPI().GetByID())
	app.Server().AddRoute("/api/transactions/{id}", "PUT", transaction.REST().UpdateByID, transaction.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/transactions/{id}", "PATCH", transaction.REST().PartiallyUpdateByID, transaction.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/transactions/{id}", "DELETE", transaction.REST().DeleteByID, transaction.OpenAPI().DeleteByID())

	app.Server().AddRoute("/api/transaction_details", "POST", transactiondetail.REST().Create, transactiondetail.OpenAPI().Create())
	app.Server().AddRoute("/api/transaction_details", "GET", transactiondetail.REST().Get, transactiondetail.OpenAPI().Get())
	app.Server().AddRoute("/api/transaction_details/{id}", "GET", transactiondetail.REST().GetByID, transactiondetail.OpenAPI().GetByID())
	app.Server().AddRoute("/api/transaction_details/{id}", "PUT", transactiondetail.REST().UpdateByID, transactiondetail.OpenAPI().UpdateByID())
	app.Server().AddRoute("/api/transaction_details/{id}", "PATCH", transactiondetail.REST().PartiallyUpdateByID, transactiondetail.OpenAPI().PartiallyUpdateByID())
	app.Server().AddRoute("/api/transaction_details/{id}", "DELETE", transactiondetail.REST().DeleteByID, transactiondetail.OpenAPI().DeleteByID())

	// AddRoute : DONT REMOVE THIS COMMENT
}
