package note

import (
	"github.com/go-pg/pg/v9"
)

type DB struct {
	*pg.DB
}

type Datastore interface {
	GetNotes() ([]*Note, error)
	GetNote(id int) (*Note, error)
	CreateNote(note *Note) (*Note, error)
	DeleteNote(id int) error
}

func (db *DB) GetNotes() ([]*Note, error) {
	var notes []*Note
	err := db.Model(&notes).Select()
	return notes, err
}

func (db *DB) GetNote(id int) (*Note, error) {
	note := &Note{Id: id}
	err := db.Select(note)
	return note, err
}

func (db *DB) CreateNote(note *Note) (*Note, error) {
	err := db.Insert(note)
	return note, err
}

func (db *DB) DeleteNote(id int) error {
	note := &Note{Id: id}
	err := db.Delete(note)
	return err
}
