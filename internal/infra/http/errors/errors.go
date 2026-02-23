package errors

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	_errors "github.com/natasha-m-oliveira/clean-architecture-go/internal/app/errors"
	"github.com/natasha-m-oliveira/clean-architecture-go/internal/infra/http/render"
)

func WriteError(w http.ResponseWriter, err error) {
	var unmarshalErr *json.UnmarshalTypeError
	var validationErr validator.ValidationErrors
	var customErr _errors.CustomError

	switch {
	case errors.As(err, &unmarshalErr):
		render.NewResponse(
			buildJSONTypeErrorResponse(unmarshalErr),
			http.StatusBadRequest,
		).Send(w)
		return

	case errors.As(err, &validationErr):
		render.NewResponse(
			buildValidationErrorResponse(validationErr),
			http.StatusBadRequest,
		).Send(w)
		return

	case errors.As(err, &customErr):
		render.NewResponse(
			map[string]any{"message": customErr.Error()},
			customErr.StatusCode(),
		).Send(w)
		return

	default:
		render.NewResponse(
			map[string]any{"message": "Internal Server Error"},
			http.StatusInternalServerError,
		).Send(w)
	}
}

func buildJSONTypeErrorResponse(err *json.UnmarshalTypeError) map[string]any {
	return map[string]any{
		"message": "Invalid request.",
		"errors": map[string]string{
			err.Field: "invalid type, expected " + err.Type.String(),
		},
	}
}

func buildValidationErrorResponse(validationErrs validator.ValidationErrors) map[string]any {
	formatted := make(map[string]string)

	for _, fieldErr := range validationErrs {
		field := fieldErr.Field()
		message := buildValidationMessage(fieldErr)

		formatted[field] = message
	}

	return map[string]any{
		"message": "Invalid request.",
		"errors":  formatted,
	}
}

func buildValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valid email"
	case "min":
		return "must have at least " + err.Param() + " characters"
	case "max":
		return "must have at most " + err.Param() + " characters"
	default:
		return "is invalid"
	}
}
