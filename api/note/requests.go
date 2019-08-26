package note

import (
	"errors"
	"net/http"
)

type NoteRequest struct {
	*Note
}

func (note *NoteRequest) Bind(r *http.Request) error {
	if note.Note == nil {
		return errors.New("Missing required JSON attributes.")
	}
	return note.Note.Validate()
}
