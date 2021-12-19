package slices

import (
	"constraints"
)

// Sum returns the sum of the elements in the slice.
//
// If s is a slice of strings(or has string as underlying type), it returns concatenated string.
func Sum[T constraints.Ordered](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}

	return sum
}

// Sum returns the sum of the elements in the slice after applying f to each element.
//
// If returned type R is a string(or has string as underlying type), it returns concatenated string.
func SumFn[T any, R constraints.Ordered](s []T, f func(T) R) R {
	var sum R
	for _, v := range s {
		sum += f(v)
	}

	return sum
}
