package errors

import "net/http"

type ProductNotFound struct {
	message string
}

func NewProductNotFound() CustomError {
	return ProductNotFound{
		message: "product not found",
	}
}

func (e ProductNotFound) Error() string {
	return e.message
}

func (e ProductNotFound) StatusCode() int {
	return http.StatusNotFound
}
