package company

type Datastore interface {
	GetCompanies() (*[]Company, error)
	GetCompany(id int) (*Company, error)
}

func (db *DB) GetCompanies() (*[]Company, error) {
	companies := new([]Company)
	err := db.Model(companies).Select()
	return companies, err
}

func (db *DB) GetCompany(id int) (*Company, error) {
	company := &Company{Id: id}
	err := db.Select(company)
	return company, err
}
