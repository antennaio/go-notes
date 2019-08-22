package company

import (
	"context"
	"time"

	"github.com/gosimple/slug"
	"github.com/go-ozzo/ozzo-validation"
)

type Company struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c Company) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 50)),
	)
}

func (c *Company) BeforeInsert(ctx context.Context) (context.Context, error) {
	c.Slug = slug.Make(c.Name)
	c.UpdatedAt = time.Now()
	return ctx, nil
}

func (c *Company) BeforeUpdate(ctx context.Context) (context.Context, error) {
	c.UpdatedAt = time.Now()
	return ctx, nil
}
