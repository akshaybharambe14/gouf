package slices

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args[T, R any] struct {
		f func(T) R
		s []T
	}

	// test predicate function
	f := func(t int) int {
		return t * 2
	}

	tests := []struct {
		name string
		args args[int, int]
		want []int
	}{
		{
			name: "map returns correct values",
			args: args[int, int]{f: f, s: []int{2, 4, 6, 8, 10}},
			want: []int{4, 8, 12, 16, 20},
		},
		{
			name: "empty slice",
			args: args[int, int]{f: f, s: nil},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
