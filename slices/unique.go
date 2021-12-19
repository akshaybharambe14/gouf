package slices

// Unique returns a new slice containing the unique elements of s.
func Unique[T comparable](s []T) []T {
	count := len(s)
	if count == 0 {
		return s
	}

	visited := make(map[T]struct{})
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		if _, ok := visited[s[i]]; ok {
			continue
		} else {
			result = append(result, s[i])
			visited[s[i]] = struct{}{}
		}
	}

	return result
}

func UniqueFn[T any, R comparable](s []T, f func(T) R) []T {
	count := len(s)
	if count == 0 {
		return s
	}

	visited := make(map[R]struct{})
	result := make([]T, 0, count)

	for i := 0; i < count; i++ {
		key := f(s[i])
		if _, ok := visited[key]; ok {
			continue
		} else {
			result = append(result, s[i])
			visited[key] = struct{}{}
		}
	}

	return result
}

// Unique returns a new slice containing the unique elements of s.
func UniqueInPlace[T comparable](s []T) []T {
	count := len(s)
	if count == 0 {
		return s
	}

	uniqueIndex := 0
	visited := make(map[T]struct{})
	for i := 0; i < count; i++ {
		if _, ok := visited[s[i]]; ok {
			continue
		} else {
			s[uniqueIndex] = s[i]
			visited[s[i]] = struct{}{}
			uniqueIndex++
		}
	}
	return s[:uniqueIndex]
}

// Unique returns a new slice containing the unique elements of s.
func UniqueFnInPlace[T any, R comparable](s []T, f func(T) R) []T {
	count := len(s)
	if count == 0 {
		return s
	}

	uniqueIndex := 0
	visited := make(map[R]struct{})
	for i := 0; i < count; i++ {
		key := f(s[i])
		if _, ok := visited[key]; ok {
			continue
		} else {
			s[uniqueIndex] = s[i]
			visited[key] = struct{}{}
			uniqueIndex++
		}
	}
	return s[:uniqueIndex]
}
