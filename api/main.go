package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-pg/pg/v9"

	"github.com/antennaio/go-notes/api/auth"
	"github.com/antennaio/go-notes/api/note"
	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/db"
	"github.com/antennaio/go-notes/lib/env"
)

func init() {
	env.LoadEnv()
}

// App represents the application
type App struct {
	Router *chi.Mux
	DB     *pg.DB
}

// Initialize sets up the database connection and router
func (a *App) Initialize(dbName, dbUser, dpPassword string) {
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

	a.DB = db.Connection(dbName, dbUser, dpPassword)

	// Public routes
	router.Group(func(router chi.Router) {
		router.Mount("/auth", auth.Routes(a.DB))
	})

	// Protected routes
	router.Route("/v1", func(router chi.Router) {
		tokenAuth := auth.TokenAuth()
		ds := &user.Datastore{Pg: a.DB}
		userContext := &auth.UserContext{Ds: ds}

		router.Group(func(router chi.Router) {
			router.Use(tokenAuth.Verifier())
			router.Use(tokenAuth.Authenticator())
			router.Use(userContext.Handler)
			router.Mount("/note", note.Routes(a.DB))
		})
	})

	a.Router = router
}

// Serve serves the app on the specified port
func (a *App) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func main() {
	app := App{}
	app.Initialize(
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_DB_USER"),
		os.Getenv("POSTGRES_DB_PASSWORD"))

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":8080"
	}

	app.Serve(port)
}
