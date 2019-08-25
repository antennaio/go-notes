package user

import (
	"github.com/go-pg/pg/v9"
)

type DB struct {
	*pg.DB
}

type UserDatastore interface {
	GetUserByEmail(email string) (*User, error)
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	user := new(User)
	err := db.Model(user).Where("email = ?", email).Select()
	return user, err
}
