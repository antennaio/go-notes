package tests

import (
	"os"

	"github.com/antennaio/go-notes/api/app"
	runner "github.com/antennaio/go-notes/migrations"
	"github.com/go-pg/pg/v9"
)

func makeApp() app.App {
	var a app.App
	a.Initialize(
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_TEST_DB_NAME"),
		os.Getenv("POSTGRES_DB_USER"),
		os.Getenv("POSTGRES_DB_PASSWORD"),
		false,
		false)

	return a
}

func migrateUp(db *pg.DB) {
	if runner.FirstRun(db) {
		_, _, errInit := runner.Run(db, "init")
		handleError(errInit)
	}

	_, _, errUp := runner.Run(db, "up")
	handleError(errUp)
}

func migrateDown(db *pg.DB) {
	_, _, err := runner.Run(db, "reset")
	handleError(err)
}
