package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/antennaio/go-notes/api/auth"
	"github.com/antennaio/go-notes/api/note"
	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/db"
	"github.com/antennaio/go-notes/lib/env"
)

func init() {
	env.LoadEnv()
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	log, ok := os.LookupEnv("LOG_REQUESTS")
	if ok && log == "true" {
		router.Use(middleware.Logger)
	}

	pgDb := db.Connection()

	// Public routes
	router.Group(func(router chi.Router) {
		router.Mount("/auth", auth.Routes(pgDb))
	})

	// Protected routes
	router.Route("/v1", func(router chi.Router) {
		tokenAuth := auth.TokenAuth()
		ds := &user.Datastore{pgDb}
		userContext := &auth.UserContext{ds}

		router.Group(func(router chi.Router) {
			router.Use(tokenAuth.Verifier())
			router.Use(tokenAuth.Authenticator())
			router.Use(userContext.Handler)
			router.Mount("/note", note.Routes(pgDb))
		})
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
