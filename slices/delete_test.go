package slices

import (
	"reflect"
	"testing"
)

func TestDelete(t *testing.T) {
	type args[T any] struct {
		s []T
		i int
	}
	tests := []struct {
		name        string
		args        args[int]
		wantResult  []int
		wantDeleted bool
	}{
		{
			name:        "delete from empty slice",
			args:        args[int]{s: []int{}, i: 0},
			wantResult:  []int{},
			wantDeleted: false,
		},
		{
			name:        "delete out of bound index",
			args:        args[int]{s: []int{1, 2}, i: 5},
			wantResult:  []int{1, 2},
			wantDeleted: false,
		},
		{
			name:        "delete negative index",
			args:        args[int]{s: []int{1, 2}, i: -5},
			wantResult:  []int{1, 2},
			wantDeleted: false,
		},
		{
			name:        "delete valid index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 2},
			wantResult:  []int{1, 2, 1, 2, 3},
			wantDeleted: true,
		},
		{
			name:        "delete last index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 5},
			wantResult:  []int{1, 2, 3, 1, 2},
			wantDeleted: true,
		},
		{
			name:        "delete first index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 0},
			wantResult:  []int{2, 3, 1, 2, 3},
			wantDeleted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Delete(tt.args.s, tt.args.i)
			if !reflect.DeepEqual(got, tt.wantResult) {
				t.Errorf("Delete() got = %v, want %v", got, tt.wantResult)
			}
			if got1 != tt.wantDeleted {
				t.Errorf("Delete() got1 = %v, want %v", got1, tt.wantDeleted)
			}
		})
	}
}

func TestDeleteUnordered(t *testing.T) {
	type args[T any] struct {
		s []T
		i int
	}
	tests := []struct {
		name        string
		args        args[int]
		wantResult  []int
		wantDeleted bool
	}{
		{
			name:        "delete from empty slice",
			args:        args[int]{s: []int{}, i: 0},
			wantResult:  []int{},
			wantDeleted: false,
		},
		{
			name:        "delete out of bound index",
			args:        args[int]{s: []int{1, 2}, i: 5},
			wantResult:  []int{1, 2},
			wantDeleted: false,
		},
		{
			name:        "delete negative index",
			args:        args[int]{s: []int{1, 2}, i: -5},
			wantResult:  []int{1, 2},
			wantDeleted: false,
		},
		{
			name:        "delete valid index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 2},
			wantResult:  []int{1, 2, 3, 1, 2},
			wantDeleted: true,
		},
		{
			name:        "delete last index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 5},
			wantResult:  []int{1, 2, 3, 1, 2},
			wantDeleted: true,
		},
		{
			name:        "delete first index",
			args:        args[int]{s: []int{1, 2, 3, 1, 2, 3}, i: 0},
			wantResult:  []int{3, 2, 3, 1, 2},
			wantDeleted: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := DeleteUnordered(tt.args.s, tt.args.i)
			if !reflect.DeepEqual(got, tt.wantResult) {
				t.Errorf("DeleteUnordered() got = %v, want %v", got, tt.wantResult)
			}
			if got1 != tt.wantDeleted {
				t.Errorf("DeleteUnordered() got1 = %v, want %v", got1, tt.wantDeleted)
			}
		})
	}
}

// TODO: check for errors, currently tests fail to run this test.
// func TestDeletePlaces(t *testing.T) {
// 	type args[T any] struct {
// 		s []T
// 		x []int
// 	}
// 	tests := []struct {
// 		name        string
// 		args        args[int]
// 		wantResult  []int
// 		wantDeleted bool
// 	}{
// 		{
// 			name:        "delete from empty slice",
// 			args:        args[int]{s: nil, x: nil},
// 			wantResult:  nil,
// 			wantDeleted: false,
// 		},
// 		{
// 			name:        "delete from empty indices",
// 			args:        args[int]{s: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, x: []int{}},
// 			wantResult:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
// 			wantDeleted: false,
// 		},
// 		{
// 			name:        "delete from valid indices",
// 			args:        args[int]{s: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, x: []int{0, 2, 4, 6, 8}},
// 			wantResult:  []int{1, 3, 5, 7, 9},
// 			wantDeleted: true,
// 		},
// 		{
// 			name:        "delete from invalid indices",
// 			args:        args[int]{s: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, x: []int{0, 2, 4, 6, 8, 10}},
// 			wantResult:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
// 			wantDeleted: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := DeletePlaces(tt.args.s, tt.args.x...)
// 			if !reflect.DeepEqual(got, tt.wantResult) {
// 				t.Errorf("DeletePlaces() got = %v, want %v", got, tt.wantResult)
// 			}
// 			if got1 != tt.wantDeleted {
// 				t.Errorf("DeletePlaces() got1 = %v, want %v", got1, tt.wantDeleted)
// 			}
// 		})
// 	}
// }

// TODO: check for errors, currently tests fail to run this test.
// func TestDeleteFn(t *testing.T) {
// 	type args[T any] struct {
// 		s []T
// 		f func(T) bool
// 	}

// 	f := func(i int) bool {
// 		return i%2 == 0
// 	}

// 	tests := []struct {
// 		name string
// 		args args[int]
// 		want []int
// 	}{
// 		{
// 			name: "predicate returns true for all elements",
// 			args: args[int]{s: []int{2, 4, 6, 8, 10}, f: f},
// 			want: []int{},
// 		},
// 		{
// 			name: "predicate returns false for all elements",
// 			args: args[int]{s: []int{1, 3, 5, 7, 9}, f: f},
// 			want: []int{1, 3, 5, 7, 9},
// 		},
// 		{
// 			name: "delete from empty slice",
// 			args: args[int]{s: []int{}, f: f},
// 			want: []int{},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := DeleteFn(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("DeleteFn() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
