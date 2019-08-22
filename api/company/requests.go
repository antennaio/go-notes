package company

import (
	"errors"
	"net/http"
)

type CompanyRequest struct {
	*Company
}

func (company *CompanyRequest) Bind(r *http.Request) error {
	if company.Company == nil {
		return errors.New("Missing required JSON attributes.")
	}
	return company.Company.Validate()
}
