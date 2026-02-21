package repositories

import "github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"

type ProductsRepository interface {
	Create(product *entities.Product) error
	FindById(id string) (*entities.Product, error)
	FindByName(name string) (*entities.Product, error)
	FindAll() ([]entities.Product, error)
	Save(product *entities.Product) error
	DeleteById(id string) error
}
