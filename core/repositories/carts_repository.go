package repositories

import "github.com/natasha-m-oliveira/clean-architecture-go/core/entities"

type CartsRepository interface {
	Create(cart *entities.Cart) error
	FindById(id string) (*entities.Cart, error)
	Save(cart *entities.Cart) error
	DeleteById(id string) error
}
