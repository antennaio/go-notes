package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table companies...")
		_, err := db.Exec(`
			CREATE TABLE companies (
				id SERIAL PRIMARY KEY,
				slug VARCHAR,
				name VARCHAR,
				updated_at TIMESTAMP DEFAULT CURRENT_DATE
			)
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table companies...")
		_, err := db.Exec(`DROP TABLE companies`)
		return err
	})
}
