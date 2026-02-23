package render

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	statusCode int
	result     any
}

func NewResponse(result any, status int) Response {
	return Response{
		result:     result,
		statusCode: status,
	}
}

func (success Response) Send(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(success.statusCode)
	json.NewEncoder(writer).Encode(success.result)
}
