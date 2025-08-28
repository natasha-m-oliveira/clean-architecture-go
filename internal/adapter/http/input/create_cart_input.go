package input

import "github.com/go-playground/validator/v10"

type CreateCartItemInput struct {
	ProductId string `json:"product_id" validate:"required,uuid"`
	Quantity  int    `json:"quantity" validate:"required,gte=1"`
}

type CreateCartInput struct {
	Items []CreateCartItemInput `json:"items" validate:"required,min=1,dive"`
}

func (i CreateCartInput) Validate() error {
	validate := validator.New()

	return validate.Struct(i)
}
