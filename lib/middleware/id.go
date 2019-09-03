package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/go-notes/lib/request"
	"github.com/antennaio/go-notes/lib/response"
)

func Id(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := request.ParamInt(r, "id")
		if err != nil {
			render.Render(w, r, response.BadRequest(err))
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
