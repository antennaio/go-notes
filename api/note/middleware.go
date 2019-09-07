package note

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/response"
)

type NoteContext struct {
	db Notes
}

func (m *NoteContext) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value("id").(int)
		user := r.Context().Value("user").(*user.User)

		note, err := m.db.GetForUser(id, user.Id)
		if err != nil {
			render.Render(w, r, response.NotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "note", note)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
