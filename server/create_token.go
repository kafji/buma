package server

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"kafji.net/buma/services"
)

type createTokenRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (s createTokenRequest) Bind(r *http.Request) error {
	errs := []error{}

	if s.Email == nil {
		errs = append(errs, errors.New("missing email"))
	}

	if s.Password == nil {
		errs = append(errs, errors.New("missing password"))
	}

	return errors.Join(errs...)
}

type createTokenResponse struct {
	Token string `json:"token"`
}

func createTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req createTokenRequest
	if err := render.Bind(r, &req); err != nil {
		msg := err.Error()
		badRequest(w, r, &msg)
		return
	}
	email := *req.Email
	password := *req.Password

	token, ok, err := services.CreateToken(ctx, getDB(ctx), getDB(ctx), email, password)
	if err != nil {
		switch err {
		case services.ErrEmptyEmail:
			msg := "email must not be empty"
			badRequest(w, r, &msg)
			return
		case services.ErrEmptyPassword:
			msg := "password must not be empty"
			badRequest(w, r, &msg)
		}
		panic(err)
	}
	if !ok {
		notFound(w, r, nil)
		return
	}

	res := createTokenResponse{token}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, &res)
}
