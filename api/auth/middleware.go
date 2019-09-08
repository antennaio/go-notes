package auth

import (
	"context"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/response"
)

type UserContext struct {
	Ds user.Users
}

func (m *UserContext) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		id := int(claims["id"].(float64))

		user, err := m.Ds.Get(id)
		if err != nil {
			render.Render(w, r, response.InternalServerError(err))
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
