package slices

import "testing"

func TestCount(t *testing.T) {
	type args[T any] struct {
		f func(T) bool
		s []T
	}

	// test count function
	f := func(t int) bool {
		return true
	}

	tests := []struct {
		name string
		args args[int]
		want int
	}{
		{
			name: "non-empty slice",
			args: args[int]{s: []int{1, 1, 1, 1, 1}, f: f},
			want: 5,
		},
		{
			name: "empty slice",
			args: args[int]{s: nil, f: f},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
