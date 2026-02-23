package mocks

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/app/repositories"
	"github.com/stretchr/testify/mock"
)

var _ repositories.CartsRepository = new(MockCartsRepository)

type MockCartsRepository struct {
	mock.Mock
}

// Create implements [repositories.CartsRepository].
func (m *MockCartsRepository) Create(cart *entities.Cart) error {
	args := m.Called(cart)
	return args.Error(0)
}

// DeleteById implements [repositories.CartsRepository].
func (m *MockCartsRepository) DeleteById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// FindById implements [repositories.CartsRepository].
func (m *MockCartsRepository) FindById(id string) (*entities.Cart, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Cart), args.Error(1)
}

// Save implements [repositories.CartsRepository].
func (m *MockCartsRepository) Save(cart *entities.Cart) error {
	args := m.Called(cart)
	return args.Error(0)
}
