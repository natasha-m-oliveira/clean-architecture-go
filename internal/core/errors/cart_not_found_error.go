package errors

type CartNotFound struct{}

var _ error = (*CartNotFound)(nil)

func (e CartNotFound) Error() string {
	return "cart not found"
}
