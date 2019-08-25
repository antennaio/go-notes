package auth

import (
	"net/http"
)

type LoginResponse struct {
	Token string `json:"token"`
}

// Pre-processing before a response is marshalled and sent across the wire
func (response *LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewLoginResponse(token string) *LoginResponse {
	resp := &LoginResponse{Token: token}
	return resp
}
