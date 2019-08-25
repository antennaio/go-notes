package company

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/goapi/lib/error"
	"github.com/antennaio/goapi/lib/request"
)

func (env *Env) getCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := env.db.GetCompanies()
	if err != nil {
		render.Render(w, r, error.InternalServerError(err))
		return
	}
	if err := render.RenderList(w, r, NewCompanyListResponse(companies)); err != nil {
		render.Render(w, r, error.InternalServerError(err))
		return
	}
}

func (env *Env) getCompany(w http.ResponseWriter, r *http.Request) {
	id, err := request.ParamInt(r, "id")
	if err != nil {
		render.Render(w, r, error.BadRequest(err))
		return
	}

	company, err := env.db.GetCompany(id)
	if err != nil {
		render.Render(w, r, error.NotFound)
		return
	}
	if err := render.Render(w, r, NewCompanyResponse(company)); err != nil {
		render.Render(w, r, error.InternalServerError(err))
		return
	}
}

func (env *Env) createCompany(w http.ResponseWriter, r *http.Request) {
	data := &CompanyRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, error.UnprocessableEntity(err))
		return
	}

	company := data.Company
	company, err := env.db.CreateCompany(company)
	if err != nil {
		render.Render(w, r, error.InternalServerError(err))
		return
	}
	if err := render.Render(w, r, NewCompanyResponse(company)); err != nil {
		render.Render(w, r, error.InternalServerError(err))
		return
	}
}
