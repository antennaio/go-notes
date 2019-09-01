package note

import (
	"github.com/go-pg/pg/v9"
)

type DB struct {
	Pg *pg.DB
}

type Notes interface {
	GetAll() ([]*Note, error)
	Get(id int) (*Note, error)
	Create(note *Note) (*Note, error)
	Update(note *Note) (*Note, error)
	Delete(id int) error
}

func (db *DB) GetAll() ([]*Note, error) {
	var notes []*Note
	err := db.Pg.Model(&notes).Select()
	return notes, err
}

func (db *DB) Get(id int) (*Note, error) {
	note := &Note{Id: id}
	err := db.Pg.Select(note)
	return note, err
}

func (db *DB) Create(note *Note) (*Note, error) {
	err := db.Pg.Insert(note)
	return note, err
}

func (db *DB) Update(note *Note) (*Note, error) {
	_, err := db.Pg.Model(note).WherePK().Update()
	return note, err
}

func (db *DB) Delete(id int) error {
	note := &Note{Id: id}
	err := db.Pg.Delete(note)
	return err
}
