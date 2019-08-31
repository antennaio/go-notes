package note

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"
)

type Env struct {
	db Datastore
}

func Routes(pgDb *pg.DB) *chi.Mux {
	db := &DB{pgDb}
	env := &Env{db}

	router := chi.NewRouter()

	router.Get("/", env.getNotes)
	router.Get("/{id}", env.getNote)
	router.Post("/", env.createNote)
	router.Put("/{id}", env.updateNote)
	router.Delete("/{id}", env.deleteNote)

	return router
}
