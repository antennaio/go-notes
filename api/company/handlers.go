package company

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/goapi/lib/request"
)

func (env *Env) getCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := env.db.GetCompanies()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	render.JSON(w, r, companies)
}

func (env *Env) getCompany(w http.ResponseWriter, r *http.Request) {
	id, err := request.ParamInt(r, "id")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	company, err := env.db.GetCompany(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	render.JSON(w, r, company)
}
