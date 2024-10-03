package mappers

import (
	"reflect"

	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infrastructure/database/prisma/utils"
	"github.com/natasha-m-oliveira/clean-architecture-go/prisma/db"
)

type PrismaCartMapper struct{}

func (m PrismaCartMapper) ToPrisma(cart entities.Cart) (status db.CartWithPrismaStatusSetParam, optional []db.CartSetParam) {
	status = db.Cart.Status.Set(db.CartStatus(cart.Status))

	totalFields := reflect.TypeOf(entities.Cart{}).NumField()
	optional = make([]db.CartSetParam, 0, totalFields-1)

	optional = append(optional,
		db.Cart.ID.SetIfPresent(&cart.Id),
		db.Cart.CreatedAt.Set(cart.CreatedAt),
		db.Cart.UpdatedAt.Set(cart.UpdatedAt))

	return status, optional
}

func (m PrismaCartMapper) ToDomain(model db.CartModel) entities.Cart {
	status := model.Status
	items := utils.MapToDomain(model.Items(), (PrismaCartItemMapper{}).ToDomain)

	optional := entities.CartOptions{
		Id:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	cart := entities.NewCart(entities.CartStatus(status), items, optional)

	return *cart
}
