package input

import "github.com/go-playground/validator/v10"

type CreateCartInput struct {
	Items []struct {
		ProductId string `json:"productId" validate:"required,min=3"`
		Quantity  int    `json:"quantity" validate:"required,gte=1"`
	} `json:"items" validate:"required,dive,required"`
}

func (i CreateCartInput) Validate() error {
	validate := validator.New()

	return validate.Struct(i)
}
