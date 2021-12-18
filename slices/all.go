package slices

// All returns true if predicate function f returns true for all elements OR if the slice is empty.
//
// If the predicate function f returns false for any element, it stops the iteration and returns false.
func All[T any](s []T, f func(T) bool) bool {
	for _, x := range s {
		if !f(x) {
			return false
		}
	}
	return true
}
