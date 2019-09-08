package note

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/go-notes/api/auth"
	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/middleware"
	"github.com/antennaio/go-notes/lib/response"
)

// NoteContextKey is a concrete type used as a key, the point is to avoid collisions between packages using context
type NoteContextKey struct{}

// NoteContext middleware depends on a datastore
type NoteContext struct {
	Ds Notes
}

// Handler injects a note into the request context
func (m *NoteContext) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(middleware.IDContextKey{}).(int)
		user := r.Context().Value(auth.UserContextKey{}).(*user.User)

		note, err := m.Ds.GetForUser(id, user.Id)
		if err != nil {
			render.Render(w, r, response.NotFound)
			return
		}

		ctx := context.WithValue(r.Context(), NoteContextKey{}, note)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
