package errors

type ProductNotFound struct{}

var _ error = (*ProductNotFound)(nil)

func (e ProductNotFound) Error() string {
	return "product not found"
}
