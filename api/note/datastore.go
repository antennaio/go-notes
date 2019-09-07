package note

import (
	"github.com/go-pg/pg/v9"
)

type DB struct {
	Pg *pg.DB
}

type Notes interface {
	GetAll() ([]*Note, error)
	GetAllForUser(userId int) ([]*Note, error)
	Get(id int) (*Note, error)
	GetForUser(id int, userId int) (*Note, error)
	Create(note *Note) (*Note, error)
	Update(note *Note) (*Note, error)
	Delete(id int) error
}

func (db *DB) GetAll() ([]*Note, error) {
	var notes []*Note
	err := db.Pg.Model(&notes).Select()
	return notes, err
}

func (db *DB) GetAllForUser(userId int) ([]*Note, error) {
	var notes []*Note
	err := db.Pg.Model(&notes).
		Where("user_id = ?", userId).
		Select()
	return notes, err
}

func (db *DB) Get(id int) (*Note, error) {
	note := &Note{Id: id}
	err := db.Pg.Select(note)
	return note, err
}

func (db *DB) GetForUser(id int, userId int) (*Note, error) {
	note := new(Note)
	err := db.Pg.Model(note).
		Where("id = ?", id).
		Where("user_id = ?", userId).
		Select()
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
