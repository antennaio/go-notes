package tests

import (
	"fmt"

	"github.com/antennaio/go-notes/api/auth"
	"github.com/antennaio/go-notes/api/note"
	"github.com/antennaio/go-notes/api/user"
	"github.com/go-pg/pg/v9"
)

type Generator struct {
	Pg *pg.DB
}

func (g *Generator) generateUser() (*user.User, string) {
	ds := user.Datastore{Pg: g.Pg}
	user := &user.User{
		FirstName: "John",
		LastName:  "Tester",
		Email:     "john@tester.com",
		Password:  "",
	}
	user, err := ds.Create(user)
	if err != nil {
		handleError(err)
	}

	tokenAuth := auth.TokenAuth()
	token := tokenAuth.EncodeToken(user)

	return user, token
}

func (g *Generator) generateNotes(u *user.User, n int) []*note.Note {
	notes := make([]*note.Note, n)
	ds := note.Datastore{Pg: g.Pg}

	for i := 0; i < n; i++ {
		note := &note.Note{
			UserId:  u.Id,
			Title:   fmt.Sprintf("Title %d", i),
			Content: fmt.Sprintf("Content %d", i),
		}
		note, err := ds.Create(note)
		if err != nil {
			handleError(err)
		}
		notes = append(notes, note)
	}

	return notes
}

func (g *Generator) truncateNotes() {
	if _, err := g.Pg.Exec("TRUNCATE TABLE notes"); err != nil {
		handleError(err)
	}
}
