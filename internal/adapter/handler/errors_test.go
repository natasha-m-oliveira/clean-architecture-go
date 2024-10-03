package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	_errors "github.com/natasha-m-oliveira/clean-architecture-go/internal/core/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandleErrors(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		statusCode int
	}{
		{
			name:       "ProductNotFound should return status not found",
			err:        _errors.ProductNotFound{},
			statusCode: http.StatusNotFound,
		},
		{
			name:       "ProductAlreadyExists should return status bad request",
			err:        _errors.ProductAlreadyExists{},
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Other error should return status internal server error",
			err:        errors.New("other error"),
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			HandleErrors(w, test.err)

			assert.Equal(t, test.statusCode, w.Result().StatusCode)
		})
	}
}
