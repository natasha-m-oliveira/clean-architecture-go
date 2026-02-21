package mappers

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/sqlc/db"
)

type SqlcCartMapper struct{}

func (m SqlcCartMapper) ToSqlc(cart entities.Cart) db.Cart {
	status := db.CartStatus(cart.Status)

	return db.Cart{
		ID:        cart.Id,
		Status:    status,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
}

func (m SqlcCartMapper) ToDomain(model db.Cart) *entities.Cart {
	optional := entities.CartOptions{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	cart := entities.NewCart(entities.CartStatus(model.Status), []entities.CartItem{}, optional)

	return cart
}
