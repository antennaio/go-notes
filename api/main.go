package main

import (
	"fmt"

	"github.com/antennaio/go-notes/api/app"
	"github.com/antennaio/go-notes/lib/env"
)

func init() {
	env.LoadEnv(".env")
}

func main() {
	a := app.App{}
	a.Initialize(
		env.Getenv("POSTGRES_DB_HOST"),
		env.Getenv("POSTGRES_DB_PORT"),
		env.Getenv("POSTGRES_DB_NAME"),
		env.Getenv("POSTGRES_DB_USER"),
		env.Getenv("POSTGRES_DB_PASSWORD"),
		env.GetenvBool("LOG_REQUESTS"),
		env.GetenvBool("LOG_QUERIES"))

	port := fmt.Sprintf(":%s", env.Getenv("PORT"))

	a.Serve(port)
}
