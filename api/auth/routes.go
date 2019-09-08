package auth

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"

	"github.com/antennaio/go-notes/api/user"
)

// Env is used to inject a datastore into request handlers
type Env struct {
	Ds user.Users
}

// Routes sets up the router
func Routes(pgDb *pg.DB) *chi.Mux {
	ds := &user.Datastore{Pg: pgDb}
	env := &Env{ds}

	router := chi.NewRouter()
	router.Post("/", env.login)
	router.Post("/register", env.register)

	return router
}
