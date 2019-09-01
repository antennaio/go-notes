package note

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
)

type Middleware struct {
	db Notes
}

type Env struct {
	db Notes
}

func Routes(pgDb *pg.DB) *chi.Mux {
	db := &DB{pgDb}
	env := &Env{db}
	middleware := &Middleware{db}

	router := chi.NewRouter()

	router.Get("/", env.getNotes)
	router.Post("/", env.createNote)

	router.Route("/{id}", func(router chi.Router) {
		router.Use(middleware.NoteContext)
		router.Get("/", env.getNote)
		router.Put("/", env.updateNote)
		router.Delete("/", env.deleteNote)
	})

	return router
}
