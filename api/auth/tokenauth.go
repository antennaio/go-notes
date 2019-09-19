package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/response"
)

var auth *Auth
var once sync.Once

type Auth struct {
	jwtAuth *jwtauth.JWTAuth
}

// EncodeToken encodes a token for a user
func (a *Auth) EncodeToken(user *user.User) string {
	_, token, _ := a.jwtAuth.Encode(jwt.MapClaims{"id": user.Id, "email": user.Email})
	return token
}

// Verifier wraps the default implementation of a Verifier from go-chi/jwtauth
func (a *Auth) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(a.jwtAuth)
}

// Authenticator is a custom implementation of an Authenticator
func (a *Auth) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())

			if err != nil || token == nil || !token.Valid {
				render.Render(w, r, response.Unauthorized(errors.New("Invalid token.")))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// TokenAuth uses the singleton pattern to initialize the auth service
func TokenAuth() *Auth {
	once.Do(func() {
		secret, ok := os.LookupEnv("APP_KEY")
		if !ok || secret == "" {
			log.Panic("App key is not set.")
		}

		auth = &Auth{jwtauth.New("HS256", []byte(secret), nil)}
	})

	return auth
}
