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

func Migrator() *migratorUtil {
	if migrator == nil {
		migrator = &migratorUtil{}
		migrator.Configure()
		if app.APP_ENV == "local" || app.IS_MAIN_SERVER {
			migrator.Run()
		}
		migrator.isConfigured = true
	}
	return migrator
}

var migrator *migratorUtil

type migratorUtil struct {
	isConfigured bool
}

func (*migratorUtil) Configure() {
	app.DB().RegisterTable("main", user.User{})
	app.DB().RegisterTable("main", medicalofficer.MedicalOfficer{})
	app.DB().RegisterTable("main", medicaltreatment.MedicalTreatment{})
	app.DB().RegisterTable("main", registration.Registration{})
	app.DB().RegisterTable("main", schedule.Schedule{})
	app.DB().RegisterTable("main", transaction.Transaction{})
	app.DB().RegisterTable("main", transactiondetail.TransactionDetail{})
	// RegisterTable : DONT REMOVE THIS COMMENT
}

func (*migratorUtil) Run() {
	tx, err := app.DB().Conn("main")
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	} else {
		err = app.DB().MigrateTable(tx, "main", app.Setting{})
	}
	if err != nil {
		app.Logger().Fatal().Err(err).Send()
	}
}
