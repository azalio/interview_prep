package array_two

import (
	"reflect"
	"testing"
)

var test = []struct {
	s    []int
	want []int
}{
	{
		[]int{-5, -3, 0, 1, 4},
		[]int{0, 1, 9, 16, 25},
	},
	{
		[]int{-5, -4, -3, -2, -1},
		[]int{1, 4, 9, 16, 25},
	},
}

func TestSortedSquares(t *testing.T) {
	//var ss []int
	for _, tt := range test {
		//copy(ss, tt.s)
		result := sortedSquares(tt.s)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("\nWant: %v, got: %v\n", tt.want, tt.s)
		}
	}
}
