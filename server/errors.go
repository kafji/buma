package server

import (
	"net/http"

	"github.com/go-chi/render"
)

type errorResponse struct {
	Error errorPayload `json:"error"`
}

type errorPayload struct {
	Message *string `json:"message"`
}

func newErrorResponse(msg *string) errorResponse {
	return errorResponse{errorPayload{msg}}
}

func badRequest(w http.ResponseWriter, r *http.Request, msg *string) {
	res := newErrorResponse(msg)
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, &res)
}

func forbidden(w http.ResponseWriter, r *http.Request, msg *string) {
	res := newErrorResponse(msg)
	render.Status(r, http.StatusForbidden)
	render.JSON(w, r, &res)
}

func notFound(w http.ResponseWriter, r *http.Request, msg *string) {
	res := newErrorResponse(msg)
	render.Status(r, http.StatusNotFound)
	render.JSON(w, r, &res)
}
