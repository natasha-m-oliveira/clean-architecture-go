package errors

type ProductAlreadyExists struct{}

var _ error = (*ProductAlreadyExists)(nil)

func (e ProductAlreadyExists) Error() string {
	return "product already exists"
}
