package numberOfGoodPairs

import "testing"

var test = []struct {
	nums []int
	want int
}{
	{
		[]int{1, 2, 3, 1, 1, 3},
		4,
	},
	{
		[]int{1, 1, 1, 1},
		6,
	},
	{
		[]int{1, 2, 3},
		0,
	},
}

func TestNumIdenticalPairs(t *testing.T) {
	for _, tt := range test {
		result := numIdenticalPairs(tt.nums)
		if result != tt.want {
			t.Errorf("func TestNumIdenticalPairs\nnums: %v\nwant: %d,"+
				" get: %d\n", tt.nums, tt.want, result)
		}
	}
}

func TestNumIdenticalPairs2(t *testing.T) {
	for _, tt := range test {
		result := numIdenticalPairs2(tt.nums)
		if result != tt.want {
			t.Errorf("func TestNumIdenticalPairs2\nnums: %v\nwant: %d,"+
				" get: %d\n", tt.nums, tt.want, result)
		}
	}
}
