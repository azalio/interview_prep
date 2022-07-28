package moveZeroes

import (
	"reflect"
	"testing"
)

var test = []struct {
	nums []int
	want []int
}{
	{
		[]int{0, 1, 0, 3, 12},
		[]int{1, 3, 12, 0, 0},
	},
	{
		[]int{0},
		[]int{0},
	},
	{
		[]int{0, 0, 1},
		[]int{1, 0, 0},
	},
}

func TestMoveZeroes(t *testing.T) {
	for _, tt := range test {
		moveZeroes(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("Want: %d, got: %d\n", tt.want, tt.nums)
		}
	}
}
