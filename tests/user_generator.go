package tests

import (
	"github.com/antennaio/go-notes/api/auth"
	"github.com/antennaio/go-notes/api/user"
	"github.com/go-pg/pg/v9"
)

func generateUser(db *pg.DB) (*user.User, string) {
	ds := user.Datastore{Pg: db}
	user := &user.User{
		FirstName: "John",
		LastName:  "Tester",
		Email:     "john@tester.com",
		Password:  "",
	}
	user, err := ds.Create(user)
	if err != nil {
		handleError(err)
	}

	tokenAuth := auth.TokenAuth()
	token := tokenAuth.EncodeToken(user)

	return user, token
}
