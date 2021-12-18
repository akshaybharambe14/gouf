package slices

import "testing"

func TestAll(t *testing.T) {
	type args[T any] struct {
		f func(T) bool
		s []T
	}

	// test predicate function
	f := func(t int) bool {
		return t > 0
	}

	tests := []struct {
		name string
		args args[int]
		want bool
	}{
		{
			name: "all elements satisfy predicate",
			args: args[int]{f: f, s: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "some elements satisfy predicate",
			args: args[int]{f: f, s: []int{1, 2, 3, 4, 5, 0}},
			want: false,
		},
		{
			name: "empty slice",
			args: args[int]{f: f, s: nil},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
