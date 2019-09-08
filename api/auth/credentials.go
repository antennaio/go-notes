package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Credentials defines a payload that is received on user login
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates the Credentials
func (c Credentials) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Password, validation.Required),
	)
}
