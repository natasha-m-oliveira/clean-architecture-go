package mocks

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/repositories"
	"github.com/stretchr/testify/mock"
)

var _ repositories.ProductsRepository = new(MockProductsRepository)

type MockProductsRepository struct {
	mock.Mock
}

// Create implements [repositories.ProductsRepository].
func (m *MockProductsRepository) Create(product *entities.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

// DeleteById implements [repositories.ProductsRepository].
func (m *MockProductsRepository) DeleteById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// FindById implements [repositories.ProductsRepository].
func (m *MockProductsRepository) FindById(id string) (*entities.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Product), args.Error(1)
}

// FindByName implements [repositories.ProductsRepository].
func (m *MockProductsRepository) FindByName(name string) (*entities.Product, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Product), args.Error(1)
}

// FindAll implements [repositories.ProductsRepository].
func (m *MockProductsRepository) FindAll() ([]entities.Product, error) {
	args := m.Called()
	return args.Get(0).([]entities.Product), args.Error(1)
}

// Save implements [repositories.ProductsRepository].
func (m *MockProductsRepository) Save(product *entities.Product) error {
	args := m.Called(product)
	return args.Error(0)
}
