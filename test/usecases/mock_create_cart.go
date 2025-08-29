package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	"github.com/natasha-m-oliveira/clean-architecture-go/core/usecases"
	"github.com/stretchr/testify/mock"
)

type MockCreateCartUseCase struct {
	mock.Mock
}

func (m *MockCreateCartUseCase) Execute(req usecases.CreateCartRequest) (*entities.Cart, error) {
	res := m.Called(req)

	if res.Get(0) == nil {
		return nil, res.Error(1)
	}

	return res.Get(0).(*entities.Cart), res.Error(1)
}
