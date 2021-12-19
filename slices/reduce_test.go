package slices

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	type args[T, R any] struct {
		initial R
		f       func(R, T) R
		s       []T
	}

	// test reducer function
	f := func(acm, t int) int {
		return acm + t
	}

	tests := []struct {
		name string
		args args[int, int]
		want int
	}{
		{
			name: "reduce with non-empty initial value",
			args: args[int, int]{initial: 10, f: f, s: []int{1, 1, 1, 1, 1}},
			want: 15,
		},
		{
			name: "reduce with empty initial value",
			args: args[int, int]{initial: 0, f: f, s: []int{1, 1, 1, 1, 1}},
			want: 5,
		},
		{
			name: "empty slice",
			args: args[int, int]{initial: 0, f: f, s: []int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.s, tt.args.initial, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
