package slices

import (
	"constraints"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	type args[T constraints.Integer | constraints.Float] struct {
		s []T
	}
	tests := []struct {
		name string
		args args[int]
		want int
	}{
		{
			name: "sum of int slice",
			args: args[int]{s: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
		{
			name: "emptry slice",
			args: args[int]{s: nil},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumStrings(t *testing.T) {
	type args[T ~string] struct {
		s []T
	}

	tests := []struct {
		name string
		want string
		args args[string]
	}{
		{
			name: "sum of int slice",
			args: args[string]{s: []string{"1", "2", "3", "4", "5"}},
			want: "12345",
		},
		{
			name: "emptry slice",
			args: args[string]{s: nil},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFn(t *testing.T) {
	type employee struct {
		salary float64
	}

	type args[T any, R constraints.Ordered] struct {
		f func(T) R
		s []T
	}

	// test reducer function
	f := func(e employee) float64 {
		return e.salary
	}

	tests := []struct {
		name string
		args args[employee, float64]
		want float64
	}{
		{
			name: "sum of all values",
			args: args[employee, float64]{s: []employee{{123.123}, {123.123}}, f: f},
			want: 246.246,
		},
		{
			name: "emptry slice",
			args: args[employee, float64]{s: nil, f: f},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumFn(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
