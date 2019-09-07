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
				user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
				slug VARCHAR,
				title VARCHAR,
				content TEXT,
				created_at TIMESTAMP DEFAULT CURRENT_DATE,
				updated_at TIMESTAMP DEFAULT CURRENT_DATE,
				deleted_at TIMESTAMP
			)
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table notes...")
		_, err := db.Exec(`DROP TABLE notes`)
		return err
	})
}
