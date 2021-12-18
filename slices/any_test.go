package slices

import "testing"

func TestAny(t *testing.T) {
	type args[T any] struct {
		f func(T) bool
		s []T
	}

	// test predicate function
	f := func(t int) bool {
		return t == 0
	}

	tests := []struct {
		name string
		args args[int]
		want bool
	}{
		{
			name: "predicate returns true for any element",
			args: args[int]{f: f, s: []int{0, 1, 2, 3, 4}},
			want: true,
		},
		{
			name: "predicate returns false for all elements",
			args: args[int]{f: f, s: []int{1, 2, 3, 4, 5}},
			want: false,
		},
		{
			name: "empty slice",
			args: args[int]{f: f, s: nil},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
