package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	statusCode int
	Error      string `json:"error"`
}

func NewError(err error, status int) *Error {
	return &Error{
		Error:      err.Error(),
		statusCode: status,
	}
}

func (err Error) Send(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(err.statusCode)
	json.NewEncoder(writer).Encode(err)
}
