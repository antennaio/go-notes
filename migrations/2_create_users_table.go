package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				first_name VARCHAR,
				last_name VARCHAR,
				email VARCHAR UNIQUE,
				password VARCHAR,
				updated_at TIMESTAMP DEFAULT CURRENT_DATE
			)
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}
