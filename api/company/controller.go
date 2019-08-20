package company

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/goapi/lib/db"
	"github.com/antennaio/goapi/lib/request"
)

func GetCompany(w http.ResponseWriter, r *http.Request) {
	db := db.Connection()

	id, err := request.ParamInt(r, "id")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	company := &Company{Id: id}
	err = db.Select(company)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	render.JSON(w, r, company)
}
