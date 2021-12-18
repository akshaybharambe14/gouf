package slices

// Find returns the index of the first occurrence of v in s and true if found. It returns false otherwise.
//
// If the type is not comparable, then use the FindFn function instead.
func Find[T comparable](s []T, v T) (int, bool) {
	for i, x := range s {
		if x == v {
			return i, true
		}
	}
	return 0, false
}

// FindFn returns the index of the first occurrence of v in s and true if true is returned by predicate function f. It returns false otherwise.
func FindFn[T any](s []T, f func(T) bool) (int, bool) {
	for i, x := range s {
		if f(x) {
			return i, true
		}
	}
	return 0, false
}
