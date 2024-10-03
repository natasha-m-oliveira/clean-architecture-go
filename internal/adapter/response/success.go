package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	statusCode int
	result     any
}

func NewSuccess(result any, status int) Success {
	return Success{
		result:     result,
		statusCode: status,
	}
}

func (success Success) Send(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(success.statusCode)
	json.NewEncoder(writer).Encode(success.result)
}
