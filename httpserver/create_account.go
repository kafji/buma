package httpserver

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	createaccount "kafji.net/buma/services/create_account"
)

type createAccountRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (s createAccountRequest) Bind(r *http.Request) error {
	errs := []error{}

	if s.Email == nil {
		errs = append(errs, errors.New("missing email"))
	}

	if s.Password == nil {
		errs = append(errs, errors.New("missing password"))
	}

	return errors.Join(errs...)
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req createAccountRequest
	if err := render.Bind(r, &req); err != nil {
		badRequest(w, r, nil)
		return
	}
	email := *req.Email
	password := *req.Password

	err := createaccount.CreateAccount(ctx, getDB(ctx), email, password)
	if err != nil {
		switch err {
		case createaccount.ErrEmptyEmail:
			msg := "email must not be empty"
			badRequest(w, r, &msg)
			return
		case createaccount.ErrEmptyPassword:
			msg := "password must not be empty"
			badRequest(w, r, &msg)
			return
		case createaccount.ErrAccountAlreadyExists:
			msg := "account already exists"
			badRequest(w, r, &msg)
			return
		}
		panic(err)
	}

	render.Status(r, http.StatusOK)
}
