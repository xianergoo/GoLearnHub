package spiral_matrix_ii

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{3, 4, 7},
	}

	for _, c := range cases {
		expected := Add(c.a, c.b)
		if expected != c.expected {
			t.Errorf("%d + %d expected be %d, but %d got", c.a, c.b, c.expected, expected)
		}
	}
}

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{3}, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{"test1", args{1}, [][]int{{1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}

func Test_generateMatrix2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{3}, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{"test1", args{1}, [][]int{{1}}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
