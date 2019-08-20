package company

type Datastore interface {
	GetCompany(id int) (*Company, error)
}

func (db *DB) GetCompany(id int) (*Company, error) {
	company := &Company{Id: id}
	err := db.Select(company)
	return company, err
}
