package note

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"

	"github.com/antennaio/go-notes/lib/middleware"
)

type Env struct {
	Ds Notes
}

func Routes(pgDb *pg.DB) *chi.Mux {
	ds := &Datastore{pgDb}
	env := &Env{ds}

	router := chi.NewRouter()

	router.Get("/", env.getNotes)
	router.Post("/", env.createNote)

	router.Route("/{id}", func(router chi.Router) {
		noteContext := &NoteContext{ds}

		router.Use(middleware.Id)
		router.Use(noteContext.Handler)
		router.Get("/", env.getNote)
		router.Put("/", env.updateNote)
		router.Delete("/", env.deleteNote)
	})

	return router
}
