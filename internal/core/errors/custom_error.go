package errors

type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

type CustomError interface {
	Error() string
}

func newCustomError(message string) CustomError {
	return &customError{
		message: message,
	}
}
