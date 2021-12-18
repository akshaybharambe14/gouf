package slices

// Count returns the number of elements in the slice for which f returns true.
func Count[T any](s []T, f func(T) bool) int {
	count := 0
	for _, v := range s {
		if f(v) {
			count++
		}
	}

	return count
}
