package containsDuplicate

import (
	"testing"
)

var test = []struct {
	nums []int
	want bool
}{
	{
		[]int{1, 2, 3, 1},
		true,
	},
	{
		[]int{1, 2, 3, 4},
		false,
	},
	{
		[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		true,
	},
}

func TestContainsDuplicate(t *testing.T) {
	for _, tt := range test {
		result := containsDuplicate(tt.nums)
		if result != tt.want {
			t.Errorf("arr: %v\nWant: %t, got: %t\n", tt.nums, tt.want, result)
		}
	}
}
