package company

import (
	"net/http"

	"github.com/go-chi/render"
)

type CompanyResponse struct {
	*Company
}

// Pre-processing before a response is marshalled and sent across the wire
func (response *CompanyResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewCompanyResponse(company *Company) *CompanyResponse {
	response := &CompanyResponse{Company: company}
	return response
}

func NewCompanyListResponse(companies []*Company) []render.Renderer {
	list := []render.Renderer{}
	for _, company := range companies {
		list = append(list, NewCompanyResponse(company))
	}
	return list
}
