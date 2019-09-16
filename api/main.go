package main

import (
	"github.com/antennaio/go-notes/api/app"
	"github.com/antennaio/go-notes/lib/env"
)

func init() {
	env.LoadEnv(".env")
}

func main() {
	a := app.App{}
	a.Initialize(
		env.Getenv("POSTGRES_DB_NAME"),
		env.Getenv("POSTGRES_DB_USER"),
		env.Getenv("POSTGRES_DB_PASSWORD"),
		env.GetenvBool("LOG_REQUESTS"),
		env.GetenvBool("LOG_QUERIES"))

	port := env.GetenvWithFallback("PORT", ":8080")

	a.Serve(port)
}
