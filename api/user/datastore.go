package user

import (
	"github.com/go-pg/pg/v9"
)

type Users interface {
	Get(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	Create(user *User) (*User, error)
}

type Datastore struct {
	Pg *pg.DB
}

func (ds *Datastore) Get(id int) (*User, error) {
	user := &User{Id: id}
	err := ds.Pg.Select(user)
	return user, err
}

func (ds *Datastore) GetByEmail(email string) (*User, error) {
	user := new(User)
	err := ds.Pg.Model(user).Where("email = ?", email).Select()
	return user, err
}

func (ds *Datastore) Create(user *User) (*User, error) {
	err := ds.Pg.Insert(user)
	return user, err
}
