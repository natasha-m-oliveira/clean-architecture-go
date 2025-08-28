package errors

type NotFound struct {
	CustomError
}

func NewCartNotFound() CustomError {
	return &NotFound{CustomError: newCustomError("cart not found")}
}

func NewProductNotFound() CustomError {
	return &NotFound{CustomError: newCustomError("product not found")}
}
