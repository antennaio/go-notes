package company

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

	router.Get("/", env.getCompanies)
	router.Get("/{id}", env.getCompany)
	router.Post("/", env.createCompany)

	return router
}
