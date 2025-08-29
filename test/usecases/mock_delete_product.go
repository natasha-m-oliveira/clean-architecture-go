package usecases

import (
	"github.com/natasha-m-oliveira/clean-architecture-go/core/usecases"
	"github.com/stretchr/testify/mock"
)

type MockDeleteProductUseCase struct {
	mock.Mock
}

func (m *MockDeleteProductUseCase) Execute(req usecases.DeleteProductRequest) error {
	res := m.Called(req)

	return res.Error(0)
}
