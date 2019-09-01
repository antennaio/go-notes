package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"

	"github.com/antennaio/goapi/lib/response"
)

func (env *Env) login(w http.ResponseWriter, r *http.Request) {
	data := &LoginRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.UnprocessableEntity(err))
		return
	}

	user, err := env.db.GetByEmail(data.Credentials.Email)
	if err != nil {
		render.Render(w, r, response.Unauthorized(errors.New("Wrong credentials provided.")))
		return
	}

	password := data.Credentials.Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		render.Render(w, r, response.Unauthorized(errors.New("Wrong credentials provided.")))
		return
	}

	tokenAuth := TokenAuth()
	token := tokenAuth.EncodeToken(user)
	if err := render.Render(w, r, NewLoginResponse(token)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}

func (env *Env) register(w http.ResponseWriter, r *http.Request) {
	data := &RegisterUserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, response.UnprocessableEntity(err))
		return
	}

	user := data.User

	_, expectedErr := env.db.GetByEmail(user.Email)
	if expectedErr == nil {
		render.Render(w, r, response.BadRequest(errors.New("User account already exists.")))
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	user, err := env.db.Create(user)
	if err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}

	tokenAuth := TokenAuth()
	token := tokenAuth.EncodeToken(user)
	if err := render.Render(w, r, NewLoginResponse(token)); err != nil {
		render.Render(w, r, response.InternalServerError(err))
		return
	}
}
