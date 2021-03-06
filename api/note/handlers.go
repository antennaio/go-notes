package note

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/go-notes/api/auth"
	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/middleware"
	"github.com/antennaio/go-notes/lib/response"
)

func (env *Env) getNotes(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(auth.UserContextKey{}).(*user.User)
	notes, err := env.Ds.GetAllForUser(user.Id)
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.RenderList(w, r, NewNoteListResponse(notes)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) getNote(w http.ResponseWriter, r *http.Request) {
	note := r.Context().Value(NoteContextKey{}).(*Note)

	if err := render.Render(w, r, NewNoteResponse(note)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) createNote(w http.ResponseWriter, r *http.Request) {
	data := &NoteRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.UnprocessableEntity(err))
		return
	}

	user := r.Context().Value(auth.UserContextKey{}).(*user.User)

	note := data.Note
	note.UserId = user.Id
	note, err := env.Ds.Create(note)
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.Render(w, r, NewNoteResponse(note)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) updateNote(w http.ResponseWriter, r *http.Request) {
	data := &NoteRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.UnprocessableEntity(err))
		return
	}

	note := r.Context().Value(NoteContextKey{}).(*Note)
	note.Title = data.Note.Title
	note.Content = data.Note.Content
	note, err := env.Ds.Update(note)
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
	if err := render.Render(w, r, NewNoteResponse(note)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) deleteNote(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.IDContextKey{}).(int)

	if err := env.Ds.Delete(id); err != nil {
		render.Render(w, r, response.NotFound)
		return
	}
	if err := render.Render(w, r, response.Success); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}
