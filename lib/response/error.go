package response

import (
	"net/http"
	"unicode"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	AppCode        int64  `json:"code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func BadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Bad request",
		ErrorText:      formatErrorMessage(err.Error()),
	}
}

func UnprocessableEntity(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Unprocessable entity",
		ErrorText:      formatErrorMessage(err.Error()),
	}
}

func InternalServerError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Internal server error",
		ErrorText:      formatErrorMessage(err.Error()),
	}
}

func Unauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 401,
		StatusText:     "Unauthorized",
		ErrorText:      formatErrorMessage(err.Error()),
	}
}

var NotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Not found"}

func formatErrorMessage(message string) string {
	lastChar := message[len(message)-1:]
	dot := string('.')
	if lastChar != dot {
		message = message + dot
	}
	for i, v := range message {
		return string(unicode.ToUpper(v)) + message[i+1:]
	}
	return ""
}
