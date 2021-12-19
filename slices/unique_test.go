package slices

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "all unique elements",
			args: args[int]{s: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "duplicate elements",
			args: args[int]{s: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "empty slice",
			args: args[int]{s: nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueInPlace(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "all unique elements",
			args: args[int]{s: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "duplicate elements",
			args: args[int]{s: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "empty slice",
			args: args[int]{s: nil},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueInPlace(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueFn(t *testing.T) {
	type args[T any, R comparable] struct {
		f func(T) R
		s []T
	}

	type testUnique struct {
		name  string
		value int
	}

	f := func(t testUnique) int {
		return t.value
	}

	tests := []struct {
		name string
		args args[testUnique, int]
		want []testUnique
	}{
		{
			name: "all unique elements",
			args: args[testUnique, int]{s: []testUnique{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}}, f: f},
			want: []testUnique{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}},
		},
		{
			name: "duplicate elements",
			args: args[testUnique, int]{s: []testUnique{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}, {value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}}, f: f},
			want: []testUnique{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}},
		},
		{
			name: "empty slice",
			args: args[testUnique, int]{s: nil, f: f},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueFn(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueFnInPlace(t *testing.T) {
	type args[T any, R comparable] struct {
		f func(T) R
		s []T
	}

	type testUnique struct {
		name  string
		value int
	}

	f := func(t testUnique) int {
		return t.value
	}

	tests := []struct {
		name string
		args args[testUnique, int]
		want []testUnique
	}{
		{
			name: "all unique elements",
			args: args[testUnique, int]{s: []testUnique{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}}, f: f},
			want: []testUnique{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}},
		},
		{
			name: "duplicate elements",
			args: args[testUnique, int]{s: []testUnique{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}, {value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}}, f: f},
			want: []testUnique{{value: 1}, {value: 2}, {value: 3}, {value: 4}, {value: 5}},
		},
		{
			name: "empty slice",
			args: args[testUnique, int]{s: nil, f: f},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueFnInPlace(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueFnInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
