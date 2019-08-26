package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table notes...")
		_, err := db.Exec(`
			CREATE TABLE notes (
				id SERIAL PRIMARY KEY,
				slug VARCHAR,
				title VARCHAR,
				content TEXT,
				updated_at TIMESTAMP DEFAULT CURRENT_DATE
			)
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table notes...")
		_, err := db.Exec(`DROP TABLE notes`)
		return err
	})
}
