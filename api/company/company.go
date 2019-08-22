package company

import (
	"context"
	"time"

	"github.com/gosimple/slug"
)

type Company struct {
	Id        int       `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
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
