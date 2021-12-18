package slices

import "testing"

func TestCheckIndices(t *testing.T) {
	type args[T any] struct {
		s []T
		x []int
	}
	tests := []struct {
		name string
		args args[int]
		want bool
	}{
		{
			name: "all indices are valid",
			args: args[int]{s: []int{1, 2, 3}, x: []int{0, 1, 2}},
			want: true,
		},
		{
			name: "some indices are not valid",
			args: args[int]{s: []int{1, 2}, x: []int{0, 1, 2}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIndices(tt.args.s, tt.args.x...); got != tt.want {
				t.Errorf("CheckIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
