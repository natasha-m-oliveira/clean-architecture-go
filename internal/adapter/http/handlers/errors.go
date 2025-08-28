package handlers

import (
	"errors"
	"net/http"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/adapter/http/response"
	_errors "github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
)

func HandleErrors(w http.ResponseWriter, err error) {
	var status int

	switch err.(type) {
	case *_errors.NotFound:
		status = http.StatusNotFound
	case *_errors.AlreadyExists:
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
		err = errors.New("an unexpected error occurred, please try again")
	}

	response.NewError(err, status).Send(w)
}
