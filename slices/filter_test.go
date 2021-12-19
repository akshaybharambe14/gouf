package slices

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[T any] struct {
		f func(T) bool
		s []T
	}

	// test predicate function
	f := func(t int) bool {
		return t%2 == 0
	}

	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "predicate returns true for all elements",
			args: args[int]{f: f, s: []int{2, 4, 6, 8, 10}},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "predicate returns false for some elements",
			args: args[int]{f: f, s: []int{1, 2, 3, 4, 5, 6}},
			want: []int{2, 4, 6},
		},
		{
			name: "predicate returns false for all elements",
			args: args[int]{f: f, s: []int{1, 3}},
			want: []int{},
		},
		{
			name: "empty slice",
			args: args[int]{f: f, s: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInPlace(t *testing.T) {
	type args[T any] struct {
		f func(T) bool
		s []T
	}

	// test predicate function
	f := func(t int) bool {
		return t%2 == 0
	}

	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "predicate returns true for all elements",
			args: args[int]{f: f, s: []int{2, 4, 6, 8, 10}},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "predicate returns false for some elements",
			args: args[int]{f: f, s: []int{1, 2, 3, 4, 5, 6}},
			want: []int{2, 4, 6},
		},
		{
			name: "predicate returns false for all elements",
			args: args[int]{f: f, s: []int{1, 3}},
			want: []int{},
		},
		{
			name: "empty slice",
			args: args[int]{f: f, s: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInPlace(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInPlaceGC(t *testing.T) {
	type args[T any] struct {
		f func(T) bool
		s []T
	}

	// test predicate function
	f := func(t int) bool {
		return t%2 == 0
	}

	tests := []struct {
		name string
		args args[int]
		want []int
	}{
		{
			name: "predicate returns true for all elements",
			args: args[int]{f: f, s: []int{2, 4, 6, 8, 10}},
			want: []int{2, 4, 6, 8, 10},
		},
		{
			name: "predicate returns false for some elements",
			args: args[int]{f: f, s: []int{1, 2, 3, 4, 5, 6}},
			want: []int{2, 4, 6},
		},
		{
			name: "predicate returns false for all elements",
			args: args[int]{f: f, s: []int{1, 3}},
			want: []int{},
		},
		{
			name: "empty slice",
			args: args[int]{f: f, s: nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterInPlaceGC(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInPlaceGC() = %v, want %v", got, tt.want)
			}
		})
	}
}
