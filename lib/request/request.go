package request

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ParamInt(r *http.Request, key string) (int, error) {
	val, err := strconv.Atoi(chi.URLParam(r, key))
	if err != nil {
		return 0, err
	}
	return val, nil
}
