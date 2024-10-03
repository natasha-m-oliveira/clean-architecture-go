package response

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedErr := errors.New("error test")
	expectedBody := `{"error":"` + expectedErr.Error() + `"}`
	w := httptest.NewRecorder()
	err := NewError(expectedErr, expectedStatusCode)
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
