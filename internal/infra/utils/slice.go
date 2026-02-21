package utils

func Map[Input any, Output any](in []Input, convert func(Input) Output) []Output {
	out := make([]Output, len(in))

	for index := range in {
		out[index] = convert(in[index])
	}

	return out
}
