package tests

import (
	"fmt"

	"github.com/antennaio/go-notes/api/auth"
	"github.com/antennaio/go-notes/api/note"
	"github.com/antennaio/go-notes/api/user"
	"github.com/go-pg/pg/v9"
)

// Generator generates database records for testing purposes
type Generator struct {
	Pg *pg.DB
}

func (g *Generator) generateUser() (*user.User, string) {
	user := &user.User{
		FirstName: "John",
		LastName:  "Tester",
		Email:     "john@tester.com",
		Password:  "",
	}
	if err := g.Pg.Insert(user); err != nil {
		handleError(err)
	}

	tokenAuth := auth.TokenAuth()
	token := tokenAuth.EncodeToken(user)

	return user, token
}

func (g *Generator) generateNotes(u *user.User, n int) []*note.Note {
	notes := make([]*note.Note, 0, n)

	for i := 1; i <= n; i++ {
		note := &note.Note{
			UserId:  u.Id,
			Title:   fmt.Sprintf("Title %d", i),
			Content: fmt.Sprintf("Content %d", i),
		}
		if err := g.Pg.Insert(note); err != nil {
			handleError(err)
		}
		notes = append(notes, note)
	}

	return notes
}

func (g *Generator) countNotes() int {
	count, err := g.Pg.Model((*note.Note)(nil)).Count()
	if err != nil {
		handleError(err)
	}
	return count
}

func (g *Generator) truncateNotes() {
	if _, err := g.Pg.Exec("TRUNCATE TABLE notes"); err != nil {
		handleError(err)
	}
}
