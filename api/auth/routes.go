package auth

import (
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v9"

	"github.com/antennaio/goapi/api/user"
)

type Env struct {
	db user.UserDatastore
}

func Routes(pgDb *pg.DB) *chi.Mux {
	db := &user.DB{pgDb}
	env := &Env{db}

	router := chi.NewRouter()
	router.Post("/", env.login)
	// router.Post("/register", env.register)

	return router
}
