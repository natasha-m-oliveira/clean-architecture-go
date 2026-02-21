package errors

type IntegrationError struct {
	message    string
	statusCode int
}

func NewIntegrationError(message string, statusCode int) CustomError {
	return IntegrationError{
		message:    message,
		statusCode: statusCode,
	}
}

func (e IntegrationError) Error() string {
	return e.message
}

func (e IntegrationError) StatusCode() int {
	return e.statusCode
}
