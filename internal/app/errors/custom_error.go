package errors

type CustomError interface {
	error
	StatusCode() int
}
