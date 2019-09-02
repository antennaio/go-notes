package note

import (
	"context"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/gosimple/slug"
)

type Note struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (n Note) Validate() error {
	return validation.ValidateStruct(&n,
		validation.Field(&n.Title, validation.Required, validation.Length(1, 100)),
		validation.Field(&n.Content, validation.Required),
	)
}

func (n *Note) BeforeInsert(ctx context.Context) (context.Context, error) {
	n.Slug = slug.Make(n.Title)
	n.UpdatedAt = time.Now()
	return ctx, nil
}

func (n *Note) BeforeUpdate(ctx context.Context) (context.Context, error) {
	n.Slug = slug.Make(n.Title)
	n.UpdatedAt = time.Now()
	return ctx, nil
}
