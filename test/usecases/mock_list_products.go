package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/usecases"
	"github.com/stretchr/testify/mock"
)

type MockListProductsUseCase struct {
	mock.Mock
}

func (m *MockListProductsUseCase) Execute(req usecases.ListProductsRequest) ([]entities.Product, error) {
	res := m.Called(req)

	return res.Get(0).([]entities.Product), res.Error(1)
}
