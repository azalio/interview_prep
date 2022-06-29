package findPivotIndex

import "testing"

var test = []struct {
	nums []int
	want int
}{
	{
		[]int{1, 7, 3, 6, 5, 6},
		3,
	},
	{
		[]int{1, 2, 3},
		-1,
	},
	{
		[]int{2, 1, -1},
		0,
	},
	{
		[]int{-1, -1, 1, 1, 0, 0},
		4,
	},
}

func TestPivotIndex(t *testing.T) {
	for _, tt := range test {
		result := pivotIndex(tt.nums)
		if result != tt.want {
			t.Errorf("Arr is: %v\nGot: %d\nwant: %d\n", tt.nums, result, tt.want)
		}
	}
}
