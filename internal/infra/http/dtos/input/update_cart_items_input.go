package input

import "github.com/go-playground/validator/v10"

type UpdateCartItemsInput struct {
	Items []struct {
		ProductId string `json:"productId" validate:"required"`
		Quantity  int    `json:"quantity" validate:"required,gte=1"`
	} `json:"items" validate:"required,dive,required"`
}

func (i UpdateCartItemsInput) Validate() error {
	validate := validator.New()

	return validate.Struct(i)
}
