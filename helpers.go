package walker

func Sink[T any](ch chan T) []T {
	out := []T{}

	for v := range ch {
		out = append(out, v)
	}

	return out
}
