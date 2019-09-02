package auth

import (
	"errors"
	"net/http"
	"os"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	"github.com/antennaio/goapi/api/user"
	"github.com/antennaio/goapi/lib/response"
)

var auth *tokenAuth
var once sync.Once

type tokenAuth struct {
	jwtAuth *jwtauth.JWTAuth
}

func (a *tokenAuth) EncodeToken(user *user.User) string {
	_, token, _ := a.jwtAuth.Encode(jwt.MapClaims{"id": user.Id, "email": user.Email})
	return token
}

func (a *tokenAuth) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(a.jwtAuth)
}

func (a *tokenAuth) Authenticator() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())

			if err != nil {
				render.Render(w, r, response.Unauthorized(errors.New("Invalid token.")))
				return
			}

			if token == nil || !token.Valid {
				render.Render(w, r, response.Unauthorized(errors.New("Invalid token.")))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Use the singleton pattern to initialise the TokenAuth service
func TokenAuth() *tokenAuth {
	once.Do(func() {
		secret, ok := os.LookupEnv("APP_KEY")
		if !ok || secret == "" {
			panic("App key is not set.")
		}

		auth = &tokenAuth{jwtauth.New("HS256", []byte(secret), nil)}
	})

	return auth
}
