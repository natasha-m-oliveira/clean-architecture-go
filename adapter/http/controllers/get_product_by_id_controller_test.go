package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/natasha-m-oliveira/clean-architecture-go/core/entities"
	usecases_test "github.com/natasha-m-oliveira/clean-architecture-go/test/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductByIdController(t *testing.T) {
	t.Run("should return 200 OK and product when found", func(t *testing.T) {
		mockUseCase := &usecases_test.MockGetProductByIdUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return(&entities.Product{
			Id:    "prod-1",
			Name:  "Product 1",
			Price: 100,
		}, nil)

		controller := NewGetProductByIdController(mockUseCase)

		req := httptest.NewRequest(http.MethodGet, "/product/prod-1", nil)
		req.SetPathValue("id", "prod-1")
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should call HandleErrors and return error status when use case returns error", func(t *testing.T) {
		mockUseCase := &usecases_test.MockGetProductByIdUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return(nil, errors.New("not found"))

		controller := NewGetProductByIdController(mockUseCase)

		req := httptest.NewRequest(http.MethodGet, "/product/prod-2", nil)
		req.SetPathValue("id", "prod-2")
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.NotEqual(t, http.StatusOK, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
}
