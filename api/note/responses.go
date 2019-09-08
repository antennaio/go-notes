package note

import (
	"net/http"

	"github.com/go-chi/render"
)

type NoteResponse struct {
	*Note
}

// Render takes care of pre-processing before a response is marshalled and sent across the wire
func (response *NoteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewNoteResponse(note *Note) *NoteResponse {
	response := &NoteResponse{Note: note}
	return response
}

func NewNoteListResponse(notes []*Note) []render.Renderer {
	list := []render.Renderer{}
	for _, note := range notes {
		list = append(list, NewNoteResponse(note))
	}
	return list
}
