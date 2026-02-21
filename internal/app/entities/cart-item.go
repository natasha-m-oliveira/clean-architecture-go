package entities

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	Id        string
	CartId    string
	ProductId string
	Quantity  int
	CreatedAt time.Time

	Product Product
}

type CartItemOptions struct {
	Id        string
	CreatedAt time.Time

	Product Product
}

func NewCartItem(cartId, productId string, quantity int, options ...CartItemOptions) *CartItem {
	var id string
	var createdAt time.Time
	var product Product

	if len(options) > 0 {
		id = options[0].Id
		createdAt = options[0].CreatedAt
		product = options[0].Product
	}

	if id == "" {
		id = uuid.New().String()
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	return &CartItem{
		Id:        id,
		CartId:    cartId,
		ProductId: productId,
		Quantity:  quantity,
		CreatedAt: createdAt,
		Product:   product,
	}
}
