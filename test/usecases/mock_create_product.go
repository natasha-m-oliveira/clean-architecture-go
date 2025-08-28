package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/core/usecases"
	"github.com/stretchr/testify/mock"
)

type MockCreateProductUseCase struct {
	mock.Mock
}

func (m *MockCreateProductUseCase) Execute(req usecases.CreateProductRequest) (*entities.Product, error) {
	res := m.Called(req)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	return res.Get(0).(*entities.Product), res.Error(1)
}
