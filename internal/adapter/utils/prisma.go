package utils

func MapToDomain[M any, D any](models []M, convert func(M) D) []D {
	result := make([]D, len(models))

	for index, model := range models {
		result[index] = convert(model)
	}

	return result
}
