package request

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ParamInt(r *http.Request, key string) (int, error) {
	val, err := strconv.Atoi(chi.URLParam(r, key))
	if err != nil {
		return 0, errors.New("Unable to parse URL parameter.")
	}
	return val, nil
}
