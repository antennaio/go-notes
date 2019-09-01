package note

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/goapi/lib/response"
)

func (env *Env) getNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := env.db.GetAll()
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
	note := r.Context().Value("note").(*Note)

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

	note := data.Note
	note, err := env.db.Create(note)
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

	note := data.Note
	note.Id = r.Context().Value("id").(int)
	note, err := env.db.Update(note)
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
	id := r.Context().Value("id").(int)

	if err := env.db.Delete(id); err != nil {
		render.Render(w, r, response.NotFound)
		return
	}
	if err := render.Render(w, r, response.Success); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}
