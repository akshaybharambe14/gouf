package slices

import "testing"

func TestFind(t *testing.T) {
	type args[T comparable] struct {
		v T
		s []T
	}

	type want struct {
		index int
		found bool
	}

	tests := []struct {
		name      string
		args      args[int]
		wantIndex int
		wantFound bool
	}{
		{
			name: "element exists in slice",
			args: args[int]{
				v: 1,
				s: []int{1, 2, 3, 4, 5},
			},
			wantIndex: 0,
			wantFound: true,
		},
		{
			name: "element does not existsin slice",
			args: args[int]{
				v: 6,
				s: []int{1, 2, 3, 4, 5},
			},
			wantIndex: 0,
			wantFound: false,
		},
		{
			name: "empty slice",
			args: args[int]{
				v: 6,
				s: nil,
			},
			wantIndex: 0,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := Find(tt.args.s, tt.args.v)
			if gotIndex != tt.wantIndex {
				t.Errorf("Find() got = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Find() got1 = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestFindFn(t *testing.T) {
	type testType struct {
		name string
	}

	f := func(t testType) bool {
		return t.name == "test"
	}

	type args[T any] struct {
		v T
		f func(T) bool
		s []T
	}

	type want struct {
		index int
		found bool
	}

	tests := []struct {
		name      string
		args      args[testType]
		wantIndex int
		wantFound bool
	}{
		{
			name: "element exists in slice",
			args: args[testType]{
				v: testType{name: "test"},
				s: []testType{{}, {name: "test"}, {name: "test2"}},
				f: f,
			},
			wantIndex: 1,
			wantFound: true,
		},
		{
			name: "element does not exist in slice",
			args: args[testType]{
				v: testType{name: "test3"},
				s: []testType{{}, {name: "test"}, {name: "test2"}},
				f: f,
			},
			wantIndex: 0,
			wantFound: false,
		},
		{
			name: "empty slice",
			args: args[testType]{
				v: testType{name: "test3"},
				s: nil,
				f: f,
			},
			wantIndex: 0,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotFound := Find(tt.args.s, tt.args.v)
			if gotIndex != tt.wantIndex {
				t.Errorf("Find() got = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Find() got1 = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
