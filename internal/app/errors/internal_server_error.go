package errors

import "net/http"

type InternalServerError struct {
	message string
}

func NewInternalServerError() CustomError {
	return InternalServerError{
		message: "something went wrong, try again later",
	}
}

func (e InternalServerError) Error() string {
	return e.message
}

func (e InternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}
