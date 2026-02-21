package errors

import "net/http"

type CartNotFound struct {
	message string
}

func NewCartNotFound() CustomError {
	return CartNotFound{
		message: "cart not found",
	}
}

func (e CartNotFound) Error() string {
	return e.message
}

func (e CartNotFound) StatusCode() int {
	return http.StatusNotFound
}
