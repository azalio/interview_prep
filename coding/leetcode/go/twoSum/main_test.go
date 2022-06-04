package twoSum

import "testing"
import "reflect"

func TestTwoSum(t *testing.T) {
	test := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}

	for _, tt := range test {
		result := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("nums: %+v\ntarget: %d\nGot: %+v, want: %+v", tt.nums, tt.target, result, tt.want)
		}
	}
}
