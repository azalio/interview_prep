package rotateArray

import (
	"reflect"
	"testing"
)

var test = []struct {
	nums []int
	k    int
	want []int
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 6, 7, 1, 2, 3, 4},
	},
	{
		[]int{-1, -100, 3, 99},
		2,
		[]int{3, 99, -1, -100},
	},
	{
		[]int{-1},
		2,
		[]int{-1},
	},
}

func TestRotate(t *testing.T) {
	for _, tt := range test {
		rotate(tt.nums, tt.k)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("Want: %v, got: %v", tt.want, tt.nums)
		}

	}
}
