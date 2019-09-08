package note

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"

	"github.com/antennaio/go-notes/lib/middleware"
)

// Env is used to inject a datastore into request handlers
type Env struct {
	Ds Notes
}

// Routes sets up the router
func Routes(pgDb *pg.DB) *chi.Mux {
	ds := &Datastore{Pg: pgDb}
	env := &Env{Ds: ds}

	router := chi.NewRouter()

	router.Get("/", env.getNotes)
	router.Post("/", env.createNote)

	router.Route("/{id}", func(router chi.Router) {
		noteContext := &NoteContext{Ds: ds}

		router.Use(middleware.ID)
		router.Use(noteContext.Handler)
		router.Get("/", env.getNote)
		router.Put("/", env.updateNote)
		router.Delete("/", env.deleteNote)
	})

	return router
}
