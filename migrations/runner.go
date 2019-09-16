package runner

import (
	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
)

// Run runs the migrations from the current directory
func Run(db *pg.DB, a ...string) (oldVersion, newVersion int64, err error) {
	return migrations.Run(db, a...)
}

// FirstRun determines if this is the very first migration run
func FirstRun(db *pg.DB) bool {
	var exists string
	_, err := db.Query(&exists, "SELECT to_regclass(?)", "gopg_migrations")
	if err == nil && exists == "gopg_migrations" {
		return false
	}
	return true
}
