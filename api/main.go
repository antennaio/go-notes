package main

import (
	"github.com/antennaio/go-notes/lib/env"
)

func init() {
	env.LoadEnv()
}

func main() {
	app := App{}
	app.Initialize(
		env.Getenv("POSTGRES_DB_NAME"),
		env.Getenv("POSTGRES_DB_USER"),
		env.Getenv("POSTGRES_DB_PASSWORD"),
		env.GetenvBool("LOG_REQUESTS"),
		env.GetenvBool("LOG_QUERIES"))

	port := env.GetenvWithFallback("PORT", ":8080")

	app.Serve(port)
}
