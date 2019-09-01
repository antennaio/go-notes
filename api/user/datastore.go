package user

import (
	"github.com/go-pg/pg/v9"
)

type DB struct {
	Pg *pg.DB
}

type Users interface {
	GetByEmail(email string) (*User, error)
	Create(user *User) (*User, error)
}

func (db *DB) GetByEmail(email string) (*User, error) {
	user := new(User)
	err := db.Pg.Model(user).Where("email = ?", email).Select()
	return user, err
}

func (db *DB) Create(user *User) (*User, error) {
	err := db.Pg.Insert(user)
	return user, err
}
