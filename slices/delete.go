package slices

// Delete removes the element at index i from s. It MUTATES s.
//
// Deleted element is replaced by zero value of T so that it can be garbage collected.
func Delete[T any](s []T, i int) ([]T, bool) {
	len := len(s)
	if i < 0 || i >= len {
		return s, false
	}

	if i < len-1 {
		copy(s[i:], s[i+1:])
	}

	var zero T
	s[len-1] = zero

	return s[:len-1], true
}

// DeleteUnordered removes the element at index i from s without preserving the order. It MUTATES s.
//
// Deleted element is replaced by zero value of T so that it can be garbage collected.
func DeleteUnordered[T any](s []T, i int) ([]T, bool) {
	count := len(s)
	if i < 0 || i >= count {
		return s, false
	}

	s[i] = s[count-1]

	var zero T
	s[count-1] = zero

	return s[:count-1], true
}

func DeleteFn[T any](s []T, f func(T) bool) []T {
	count := len(s)
	if count == 0 {
		return s
	}

	var numDeleted int
	for i := 0; i < count; i++ {
		if f(s[i]) {
			if i < count-1 {
				copy(s[i:], s[i+1:])
			}

			numDeleted++

			var zero T
			s[count-numDeleted] = zero
		}
	}

	return s[:count-numDeleted]
}

// DeletePlaces removes all elements at indices present in x from s. It MUTATES s.
// Passed indices must be unique and sorted in ascending order.
// It returns true if all elements were removed and false if none.
//
// Deleted element is replaced by zero value of T so that it can be garbage collected.
func DeletePlaces[T any](s []T, x ...int) ([]T, bool) {
	count := len(s)
	if count == 0 || len(x) == 0 {
		return s, false
	}

	for _, i := range x {
		if i < 0 || i >= count {
			return s, false
		}
	}

	var numDeleted int
	nextToBeDeleted := x[0]

	for i := 0; i < count; i++ {
		if i == nextToBeDeleted {
			if i < count-1 {
				copy(s[i:], s[i+1:])
			}

			numDeleted++

			var zero T
			s[count-numDeleted] = zero

			if len(x) == numDeleted {
				break
			}

			nextToBeDeleted = x[numDeleted] - numDeleted
		}
	}

	return s[:count-numDeleted], true
}
