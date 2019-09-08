package note

import (
	"github.com/go-pg/pg/v9"
)

type Notes interface {
	GetAll() ([]*Note, error)
	GetAllForUser(userId int) ([]*Note, error)
	Get(id int) (*Note, error)
	GetForUser(id int, userId int) (*Note, error)
	Create(note *Note) (*Note, error)
	Update(note *Note) (*Note, error)
	Delete(id int) error
}

type Datastore struct {
	Pg *pg.DB
}

func (ds *Datastore) GetAll() ([]*Note, error) {
	var notes []*Note
	err := ds.Pg.Model(&notes).Select()
	return notes, err
}

func (ds *Datastore) GetAllForUser(userId int) ([]*Note, error) {
	var notes []*Note
	err := ds.Pg.Model(&notes).
		Where("user_id = ?", userId).
		Select()
	return notes, err
}

func (ds *Datastore) Get(id int) (*Note, error) {
	note := &Note{Id: id}
	err := ds.Pg.Select(note)
	return note, err
}

func (ds *Datastore) GetForUser(id int, userId int) (*Note, error) {
	note := new(Note)
	err := ds.Pg.Model(note).
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Select()
	return note, err
}

func (ds *Datastore) Create(note *Note) (*Note, error) {
	err := ds.Pg.Insert(note)
	return note, err
}

func (ds *Datastore) Update(note *Note) (*Note, error) {
	_, err := ds.Pg.Model(note).WherePK().Update()
	return note, err
}

func (ds *Datastore) Delete(id int) error {
	note := &Note{Id: id}
	err := ds.Pg.Delete(note)
	return err
}
