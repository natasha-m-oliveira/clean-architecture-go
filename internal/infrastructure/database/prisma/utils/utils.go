package utils

import "errors"

func ParseError(err error, customError error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, customError) {
		return customError
	}

	return err
}

func MapToDomain[M any, D any](models []M, convert func(M) D) []D {
	result := make([]D, len(models))

	for index, model := range models {
		result[index] = convert(model)
	}

	return result
}
