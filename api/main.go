package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

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
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

    db := db.Connection()

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/company", company.Routes(db))
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
