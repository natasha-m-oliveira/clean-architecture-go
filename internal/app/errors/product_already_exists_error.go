package errors

import "net/http"

type ProductAlreadyExists struct {
	message string
}

func NewProductAlreadyExists() CustomError {
	return ProductAlreadyExists{
		message: "product already exists",
	}
}

func (e ProductAlreadyExists) Error() string {
	return e.message
}

func (e ProductAlreadyExists) StatusCode() int {
	return http.StatusConflict
}
