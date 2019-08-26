package company

import (
	"github.com/go-pg/pg/v9"
)

type DB struct {
	*pg.DB
}

type Datastore interface {
	GetCompanies() ([]*Company, error)
	GetCompany(id int) (*Company, error)
	CreateCompany(company *Company) (*Company, error)
}

func (db *DB) GetCompanies() ([]*Company, error) {
	var companies []*Company
	err := db.Model(&companies).Select()
	return companies, err
}

func (db *DB) GetCompany(id int) (*Company, error) {
	company := &Company{Id: id}
	err := db.Select(company)
	return company, err
}

func (db *DB) CreateCompany(company *Company) (*Company, error) {
	err := db.Insert(company)
	return company, err
}
