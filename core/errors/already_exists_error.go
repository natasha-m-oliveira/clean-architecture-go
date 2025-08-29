package errors

type AlreadyExists struct {
	CustomError
}

func NewProductAlreadyExists() CustomError {
	return &AlreadyExists{CustomError: newCustomError("product already exists")}
}
