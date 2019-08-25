package user

import (
	"github.com/go-pg/pg/v9"
)

type DB struct {
	*pg.DB
}

type UserDatastore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) (*User, error)
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	user := new(User)
	err := db.Model(user).Where("email = ?", email).Select()
	return user, err
}

func (db *DB) CreateUser(user *User) (*User, error) {
	err := db.Insert(user)
	return user, err
}
