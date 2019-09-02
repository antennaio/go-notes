package db

import (
	"context"
	"log"
	"os"

	"github.com/go-pg/pg/v9"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	formattedQuery, _ := q.FormattedQuery()
	log.Printf("%s\n", formattedQuery)
	return nil
}

func Connection() *pg.DB {
	db := pg.Connect(&pg.Options{
		Database: os.Getenv("POSTGRES_DB_NAME"),
		User:     os.Getenv("POSTGRES_DB_USER"),
		Password: os.Getenv("POSTGRES_DB_PASSWORD"),
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Panicf("Error: %s\n", err.Error())
	}

	log, ok := os.LookupEnv("LOG_QUERIES")
	if ok && log == "true" {
		db.AddQueryHook(dbLogger{})
	}

	return db
}
