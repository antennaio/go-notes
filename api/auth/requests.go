package auth

import (
	"errors"
	"net/http"

	"github.com/antennaio/goapi/api/user"
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

type RegisterUserRequest struct {
	*user.User
}

func (register *RegisterUserRequest) Bind(r *http.Request) error {
	if register.User == nil {
		return errors.New("Missing required JSON attributes.")
	}
	return register.User.Validate()
}
