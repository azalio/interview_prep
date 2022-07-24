package singleNumber

import "testing"

var test = []struct {
	nums []int
	want int
}{
	{
		[]int{2, 2, 1},
		1,
	},
	{
		[]int{4, 1, 2, 1, 2},
		4,
	},
	{
		[]int{1},
		1,
	},
}

func TestSingleNumber(t *testing.T) {
	for _, tt := range test {
		result := singleNumber(tt.nums)
		if result != tt.want {
			t.Errorf("Arr: %v\nWant: %d, got: %d\n", tt.nums, tt.want, result)
		}
	}
}