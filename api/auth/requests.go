package auth

import (
	"errors"
	"net/http"
)

type LoginRequest struct {
	*Credentials
}

func (login *LoginRequest) Bind(r *http.Request) error {
	if login.Credentials == nil {
		return errors.New("Missing required JSON attributes.")
	}
	return login.Credentials.Validate()
}
