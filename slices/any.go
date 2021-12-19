package slices

// Any returns true if predicate function f returns true for any element.
func Any[T any](s []T, f func(T) bool) bool {
	for _, x := range s {
		if f(x) {
			return true
		}
	}
	return false
}
