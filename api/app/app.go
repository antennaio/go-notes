package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-pg/pg/v9"

	"github.com/antennaio/go-notes/api/auth"
	"github.com/antennaio/go-notes/api/note"
	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/db"
)

// App represents the application
type App struct {
	Router *chi.Mux
	Pg     *pg.DB
}

// Initialize sets up the database connection and router
func (a *App) Initialize(dbHost, dbPort, dbName, dbUser, dpPassword string, logRequests, logQueries bool) {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	if logRequests {
		router.Use(middleware.Logger)
	}

	a.Pg = db.Connection(dbHost, dbPort, dbName, dbUser, dpPassword, logQueries)

	// Public routes
	router.Group(func(router chi.Router) {
		router.Mount("/auth", auth.Routes(a.Pg))
	})

	// Protected routes
	router.Route("/v1", func(router chi.Router) {
		tokenAuth := auth.TokenAuth()
		ds := &user.Datastore{Pg: a.Pg}
		userContext := &auth.UserContext{Ds: ds}

		router.Group(func(router chi.Router) {
			router.Use(tokenAuth.Verifier())
			router.Use(tokenAuth.Authenticator())
			router.Use(userContext.Handler)
			router.Mount("/note", note.Routes(a.Pg))
		})
	})

	a.Router = router
}

// Serve serves the app on the specified port
func (a *App) Serve(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
