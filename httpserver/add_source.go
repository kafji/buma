package httpserver

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	addsource "kafji.net/buma/services/add_source"
)

type addSourceRequest struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

func (s addSourceRequest) Bind(r *http.Request) error {
	errs := []error{}

	if s.Name == nil {
		errs = append(errs, errors.New("missing name"))
	}

	if s.URL == nil {
		errs = append(errs, errors.New("missing url"))
	}

	return errors.Join(errs...)
}

func addSourceHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req addSourceRequest
	if err := render.Bind(r, &req); err != nil {
		badRequest(w, r, nil)
		return
	}
	name := *req.Name
	url := *req.URL

	_, err := addsource.AddSource(ctx, getDB(ctx), getUserID(ctx), name, url)
	if err != nil {
		switch err {
		case addsource.ErrSourceAlreadyExists:
			badRequest(w, r, nil)
			return
		case addsource.ErrEmptySourceName:
			badRequest(w, r, nil)
			return
		case addsource.ErrEmptySourceURL:
			badRequest(w, r, nil)
			return
		}
		panic(err)
	}
}
