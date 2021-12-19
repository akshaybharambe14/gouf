package slices

// Map returns a new slice containing the results of applying mapper f to each element in s.
func Map[T, R any](s []T, f func(T) R) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(v)
	}

	return result
}
