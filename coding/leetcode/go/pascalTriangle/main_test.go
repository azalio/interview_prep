package pascalTriangle

import (
	"reflect"
	"testing"
)

var test = []struct {
	numRows int
	want    [][]int
}{
	{
		5,
		[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
	},
	{
		1,
		[][]int{{1}},
	},
}

func TestGenerate(t *testing.T) {
	for _, tt := range test {
		result := generate(tt.numRows)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Got: %+v, want: %+v", result, tt.want)
		}
	}
}
