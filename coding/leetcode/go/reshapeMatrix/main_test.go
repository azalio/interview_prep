package reshapeMatrix

import (
	"reflect"
	"testing"
)

func TestMatrixReshape(t *testing.T) {
	test := []struct {
		mat  [][]int
		r    int
		c    int
		want [][]int
	}{
		{
			[][]int{{1, 2}, {3, 4}},
			1,
			4,
			[][]int{{1, 2, 3, 4}},
		},
		{
			[][]int{{1, 2}, {3, 4}},
			2,
			4,
			[][]int{{1, 2}, {3, 4}},
		},
		{
			[][]int{{1, 2}, {3, 4}},
			4,
			1,
			[][]int{{1}, {2}, {3}, {4}},
		},
		{
			[][]int{{1, 2}},
			1,
			1,
			[][]int{{1, 2}},
		},
	}

	for _, tt := range test {
		result := matrixReshape(tt.mat, tt.r, tt.c)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Got: %+v, want: %+v", result, tt.want)
		}
	}
}
