package mappers

import (
	"reflect"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

type PrismaCartItemMapper struct{}

func (m PrismaCartItemMapper) ToPrisma(cartItem entities.CartItem) (quantity db.CartItemWithPrismaQuantitySetParam, cart db.CartItemWithPrismaCartSetParam, product db.CartItemWithPrismaProductSetParam, optional []db.CartItemSetParam) {
	quantity = db.CartItem.Quantity.Set(cartItem.Quantity)
	cart = db.CartItem.Cart.Link(db.Cart.ID.Equals(cartItem.CartId))
	product = db.CartItem.Product.Link(db.Product.ID.Equals(cartItem.ProductId))

	totalFields := reflect.TypeOf(entities.CartItem{}).NumField()
	optional = make([]db.CartItemSetParam, 0, totalFields-3)

	optional = append(optional, db.CartItem.ID.SetIfPresent(&cartItem.Id),
		db.CartItem.CreatedAt.Set(cartItem.CreatedAt))

	return quantity, cart, product, optional
}

func (m PrismaCartItemMapper) ToDomain(model db.CartItemModel) entities.CartItem {
	cartId := model.CartID
	productId := model.ProductID
	quantity := model.Quantity

	optional := entities.CartItemOptions{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
	}

	cart := entities.NewCartItem(cartId, productId, quantity, optional)

	return *cart
}
