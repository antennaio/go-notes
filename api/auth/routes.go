package auth

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"

	"github.com/antennaio/go-notes/api/user"
)

type Env struct {
	Ds user.Users
}

func Routes(pgDb *pg.DB) *chi.Mux {
	ds := &user.Datastore{pgDb}
	env := &Env{ds}

	router := chi.NewRouter()
	router.Post("/", env.login)
	router.Post("/register", env.register)

	return router
}
