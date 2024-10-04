package input

import "github.com/go-playground/validator/v10"

type CreateProductInput struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"min=3,max=20"`
	Price       int    `json:"price" validate:"required,gte=0"`
	Discount    int    `json:"discount" validate:"gte=0,discount_lte_price"`
}

func (i CreateProductInput) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("discount_lte_price", discountLtePrice)

	return validate.Struct(i)
}

func discountLtePrice(fl validator.FieldLevel) bool {
	input := fl.Parent().Interface().(CreateProductInput)
	return input.Discount <= input.Price
}
