package mappers

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/sqlc/db"
)

type SqlcCartItemMapper struct{}

func (m SqlcCartItemMapper) ToSqlc(cartItem entities.CartItem) db.CartItem {

	return db.CartItem{
		ID:        cartItem.Id,
		CartID:    cartItem.CartId,
		ProductID: cartItem.ProductId,
		Quantity:  int16(cartItem.Quantity),
		CreatedAt: cartItem.CreatedAt,
	}
}

func (m SqlcCartItemMapper) ToDomain(model db.CartItem) *entities.CartItem {
	optional := entities.CartItemOptions{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
	}

	cartItem := entities.NewCartItem(model.CartID, model.ProductID, int(model.Quantity), optional)

	return cartItem
}
