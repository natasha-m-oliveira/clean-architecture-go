package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	usecases_test "github.com/natasha-m-oliveira/clean-architecture-go/test/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteProductController(t *testing.T) {
	t.Run("should return 204 No Content when product is deleted successfully", func(t *testing.T) {
		mockUseCase := &usecases_test.MockDeleteProductUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return(nil)

		controller := NewDeleteProductController(mockUseCase)

		req := httptest.NewRequest(http.MethodDelete, "/product/123", nil)
		req.SetPathValue("id", "123")
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should call HandleErrors and return error status when use case returns error", func(t *testing.T) {
		mockUseCase := &usecases_test.MockDeleteProductUseCase{}
		mockUseCase.On("Execute", mock.Anything).Return(errors.New("delete error"))

		controller := NewDeleteProductController(mockUseCase)

		req := httptest.NewRequest(http.MethodDelete, "/product/456", nil)
		req.SetPathValue("id", "456")
		w := httptest.NewRecorder()

		controller.Execute(w, req)

		resp := w.Result()
		defer resp.Body.Close()
		assert.NotEqual(t, http.StatusNoContent, resp.StatusCode)
		mockUseCase.AssertExpectations(t)
	})
}
