package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/antennaio/goapi/api/auth"
	"github.com/antennaio/goapi/api/company"
	"github.com/antennaio/goapi/lib/db"
	"github.com/antennaio/goapi/lib/env"
)

func init() {
	env.LoadEnv()
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	db := db.Connection()
	tokenAuth := auth.TokenAuth()

	// Public routes
	router.Group(func(router chi.Router) {
		router.Mount("/auth", auth.Routes(db))
	})

	// Protected routes
	router.Group(func(router chi.Router) {
		router.Use(tokenAuth.Verifier())
		router.Use(tokenAuth.Authenticator())
		router.Mount("/company", company.Routes(db))
	})

	return router
}

func main() {
	router := Routes()

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = ":8080"
	}

	log.Fatal(http.ListenAndServe(port, router))
}
