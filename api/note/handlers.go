package note

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/antennaio/goapi/lib/response"
	"github.com/antennaio/goapi/lib/request"
)

func (env *Env) getNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := env.db.GetNotes()
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
	id, err := request.ParamInt(r, "id")
	if err != nil {
		render.Render(w, r, response.BadRequest(err))
		return
	}

	note, err := env.db.GetNote(id)
	if err != nil {
		render.Render(w, r, response.NotFound)
		return
	}
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
	note, err := env.db.CreateNote(note)
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
	id, err := request.ParamInt(r, "id")
	if err != nil {
		render.Render(w, r, response.BadRequest(err))
		return
	}

	if err := env.db.DeleteNote(id); err != nil {
		render.Render(w, r, response.NotFound)
		return
	}
	if err := render.Render(w, r, response.Success); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}
