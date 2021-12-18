package slices

// Reduce reduces the slice to a single value by applying the reducer f to each element in the slice.
//
// It returns "initial" if the slice is empty.
func Reduce[T, R any](s []T, initial R, f func(R, T) R) R {
	acm := initial
	for _, v := range s {
		acm = f(acm, v)
	}
	return acm
}
