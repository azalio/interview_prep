package removeElement

import (
	"reflect"
	"testing"
)

var test = []struct {
	nums     []int
	val      int
	wantNums []int
	want     int
}{
	{
		[]int{1},
		0,
		[]int{1},
		0,
	},
	{
		[]int{3, 2, 2, 3},
		3,
		[]int{2, 2},
		2,
	},
	{
		[]int{0, 1, 2, 2, 3, 0, 4, 2},
		2,
		[]int{0, 1, 4, 0, 3},
		5,
	},
}

func TestRemoveElement(t *testing.T) {
	for _, tt := range test {
		result := removeElement(tt.nums, tt.val)
		if result != tt.want && !reflect.DeepEqual(tt.nums, tt.wantNums) {
			t.Errorf("want: %d, %+v, got: %d, %+v",
				tt.want, tt.wantNums, tt.val, tt.nums)
		}
	}
}
