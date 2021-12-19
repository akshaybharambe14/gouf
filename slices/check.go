package slices

// CheckIndices returns true if the indices are valid for the given slice.
func CheckIndices[T any](s []T, x ...int) bool {
	count := len(s)

	for _, i := range x {
		if i < 0 || i >= count {
			return false
		}
	}

	return true
}
