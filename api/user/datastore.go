package user

import (
	"github.com/go-pg/pg/v9"
)

type DB struct {
	Pg *pg.DB
}

type Users interface {
	Get(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) (*User, error)
}

func (db *DB) Get(id int) (*User, error) {
	user := &User{Id: id}
	err := db.Pg.Select(user)
	return user, err
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
