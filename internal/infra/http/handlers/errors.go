package handlers

import (
	"net/http"

	_errors "github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
)

func HandleErrors(w http.ResponseWriter, err error) {
	var status int

	switch e := err.(type) {
	case _errors.CustomError:
		status = e.StatusCode()
	default:
		status = http.StatusInternalServerError
	}

	render.NewError(err, status).Send(w)
}
