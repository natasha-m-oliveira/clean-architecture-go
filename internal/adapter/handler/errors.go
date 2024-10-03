package handler

import (
	"errors"
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/response"
	_errors "github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
)

func HandleErrors(w http.ResponseWriter, err error) {
	var status int

	switch {
	case errors.As(err, &_errors.CartNotFound{}),
		errors.As(err, &_errors.ProductNotFound{}):
		status = http.StatusNotFound
	case errors.As(err, &_errors.ProductAlreadyExists{}):
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}

	response.NewError(err, status).Send(w)
}
