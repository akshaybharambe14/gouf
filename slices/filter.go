package slices

// Filter returns a new slice containing all elements from s for which f returns true.
//
// Elements in result are in the same order as they are in the provided slice.
func Filter[T any](s []T, f func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
