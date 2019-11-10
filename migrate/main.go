package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-pg/pg/v9"

	"github.com/antennaio/go-notes/lib/env"
	runner "github.com/antennaio/go-notes/migrations"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.

Usage:
  go run migrations/*.go <command> [args]
`

func init() {
	env.LoadEnv(".env")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	db := pg.Connect(&pg.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_PORT")),
		Database: os.Getenv("POSTGRES_DB_NAME"),
		User:     os.Getenv("POSTGRES_DB_USER"),
		Password: os.Getenv("POSTGRES_DB_PASSWORD"),
	})

	oldVersion, newVersion, err := runner.Run(db, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
