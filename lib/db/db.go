package db

import (
	"context"
	"fmt"
	"log"

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

// Connection establishes a new database connection
func Connection(dbHost, dbPort, dbName, dbUser, dbPassword string, logQueries bool) *pg.DB {
	db := pg.Connect(&pg.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%s", dbHost, dbPort),
		Database: dbName,
		User:     dbUser,
		Password: dbPassword,
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Panicf("Error: %s\n", err.Error())
	}

	if logQueries {
		db.AddQueryHook(dbLogger{})
	}

	return db
}
