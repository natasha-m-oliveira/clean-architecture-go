package render

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	type ExpectedResponse struct {
		Test string
	}

	expectedStatusCode := http.StatusOK
	expectedResult := ExpectedResponse{Test: "test"}
	expectedBody := `{"test":"test"}`
	w := httptest.NewRecorder()
	err := NewResponse(expectedResult, expectedStatusCode)
	err.Send(w)

	result := strings.TrimSpace(w.Body.String())
	if !strings.EqualFold(result, expectedBody) {
		t.Errorf(
			"[TestCase '%s'] 'Expected: '%v' | Result: '%v'",
			"test",
			expectedBody,
			result,
		)
	}

	assert.Equal(t, expectedStatusCode, w.Result().StatusCode)
	assert.Equal(t, w.Header().Get("Content-Type"), "application/json")
}
