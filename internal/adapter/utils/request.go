package utils

import (
	"encoding/json"
	"io"
)

func DecodeBody[T any](body io.ReadCloser, dto T) (*T, error) {
	jsonBody, err := io.ReadAll(body)
	defer body.Close()
	if err != nil {
		return nil, err
	}

	var input T

	if err := json.Unmarshal(jsonBody, &input); err != nil {
		return nil, err
	}

	return &input, nil
}
