package removeDuplicatesFromArray

import "testing"

var test = []struct {
	nums         []int
	expectedNums []int
	want         int
}{
	{
		[]int{1, 1, 2},
		[]int{1, 2},
		2,
	},
	{
		[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		[]int{0, 1, 2, 3, 4},
		5,
	},
	{
		[]int{0},
		[]int{0},
		1,
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range test {
		result := removeDuplicates(tt.nums)

		if result != tt.want {
			t.Errorf("Want: %d, got %d\n", tt.want, result)
		}

		for i := 0; i < result; i++ {
			if tt.nums[i] != tt.expectedNums[i] {
				t.Errorf("Want %v, got %v\n", tt.nums, tt.expectedNums)
			}
		}

	}
}
