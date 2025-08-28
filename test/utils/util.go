package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"
)

func MakeRequestBody(t *testing.T, body interface{}) io.Reader {
	b, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}
	return bytes.NewReader(b)
}
