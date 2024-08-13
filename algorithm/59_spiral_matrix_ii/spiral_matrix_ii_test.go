package spiral_matrix_ii

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	cases := []struct {
		n        int
		expected [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{1, [][]int{{1}}},
	}

	for _, c := range cases {
		actual := generateMatrix(c.n)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("For input %d, expected %v, but got %v", c.n, c.expected, actual)
		}
	}

}
