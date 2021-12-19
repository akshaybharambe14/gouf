package slices

// Filter returns a new slice containing all elements from s for which f returns true.
//
// Elements in result are in the same order as they are in the provided slice.
func Filter[T any](s []T, f func(T) bool) []T {
	if s == nil {
		return nil
	}

	result := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterInPlace MUTATES the same slice s. Returned slice contains all elements from s for which f returns true.
// Use this over Filter to avoid new memory allocations.
//
// Elements in result are in the same order as they are in the provided slice.
func FilterInPlace[T any](s []T, f func(T) bool) []T {
	// result := make([]T, 0, len(s))
	count := len(s)
	var filterIndex int
	for i := 0; i < count; i++ {
		if f(s[i]) {
			// matched, replace
			s[filterIndex] = s[i]
			filterIndex++
		}
	}
	return s[:filterIndex]
}

// FilterInPlaceGC is same as FilterInPlace other than setting all elements for which f returned false to zero value of T so that they can be garbage collected.
// Use this over FilterInPlace to avoid memory leaks.
//
// Elements in result are in the same order as they are in the provided slice.
func FilterInPlaceGC[T any](s []T, f func(T) bool) []T {
	count := len(s)
	var filterIndex int
	for i := 0; i < count; i++ {
		if f(s[i]) {
			// matched, replace
			s[filterIndex] = s[i]
			filterIndex++
		}
	}

	for i := filterIndex; i < count; i++ {
		var zero T
		s[i] = zero
	}

	return s[:filterIndex]
}
