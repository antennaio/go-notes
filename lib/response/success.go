package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type SuccessResponse struct {
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
}

func (e *SuccessResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var Success = &SuccessResponse{HTTPStatusCode: 200, StatusText: "Success"}
