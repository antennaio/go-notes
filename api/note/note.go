package note

import (
	"context"
	"time"

	"github.com/gosimple/slug"
	"github.com/go-ozzo/ozzo-validation"
)

type Note struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c Note) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 100)),
		validation.Field(&c.Content, validation.Required),
	)
}

func (c *Note) BeforeInsert(ctx context.Context) (context.Context, error) {
	c.Slug = slug.Make(c.Title)
	c.UpdatedAt = time.Now()
	return ctx, nil
}

func (c *Note) BeforeUpdate(ctx context.Context) (context.Context, error) {
	c.Slug = slug.Make(c.Title)
	c.UpdatedAt = time.Now()
	return ctx, nil
}
