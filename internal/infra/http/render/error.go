package render

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	statusCode int
	Message    string `json:"error"`
}

func NewError(err error, status int) *Error {
	return &Error{
		Message:    err.Error(),
		statusCode: status,
	}
}

func (err Error) Send(writer http.ResponseWriter) {
	
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(err.statusCode)
	json.NewEncoder(writer).Encode(err)
}
