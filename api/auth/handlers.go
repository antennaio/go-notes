package auth

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"

	"github.com/antennaio/goapi/lib/error"
)

func (env *Env) login(w http.ResponseWriter, r *http.Request) {
	data := &LoginRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, error.UnprocessableEntity(err))
		return
	}

	user, err := env.db.GetUserByEmail(data.Credentials.Email)
	if err != nil {
		render.Render(w, r, error.Unauthorized(errors.New("Wrong credentials provided.")))
		return
	}

	password := data.Credentials.Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		render.Render(w, r, error.Unauthorized(errors.New("Wrong credentials provided.")))
		return
	}

	tokenAuth := TokenAuth()
	token := tokenAuth.EncodeToken(user)
	if err := render.Render(w, r, NewLoginResponse(token)); err != nil {
		render.Render(w, r, error.InternalServerError(err))
		return
	}
}

func (env *Env) register(w http.ResponseWriter, r *http.Request) {
	data := &RegisterUserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, error.UnprocessableEntity(err))
		return
	}

	user := data.User

	exists, _ := env.db.GetUserByEmail(user.Email)
	if exists != nil {
		render.Render(w, r, error.BadRequest(errors.New("User account already exists.")))
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	user, err := env.db.CreateUser(user)
	if err != nil {
		render.Render(w, r, error.InternalServerError(err))
		return
	}

	tokenAuth := TokenAuth()
	token := tokenAuth.EncodeToken(user)
	if err := render.Render(w, r, NewLoginResponse(token)); err != nil {
		render.Render(w, r, error.InternalServerError(err))
		return
	}
}
