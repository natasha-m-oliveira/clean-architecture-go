package entities

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          string
	Name        string
	Description string
	Image       string
	Price       int
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductOptions struct {
	Id          string
	Description string
	Image       string
	Discount    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProduct(name string, price int, options ...ProductOptions) *Product {
	var id, description, image string
	var discount int
	var createdAt, updatedAt time.Time

	if len(options) > 0 {
		id = options[0].Id
		description = options[0].Description
		image = options[0].Image
		discount = options[0].Discount
		createdAt = options[0].CreatedAt
		updatedAt = options[0].UpdatedAt
	}

	if id == "" {
		id = uuid.New().String()
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = createdAt
	}

	return &Product{
		Id:          id,
		Name:        name,
		Description: description,
		Image:       image,
		Price:       price,
		Discount:    discount,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
