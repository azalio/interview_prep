package removeElement

import (
	"sort"
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
		1,
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
		[]int{0, 0, 1, 3, 4},
		5,
	},
}

func TestRemoveElement(t *testing.T) {
	for _, tt := range test {
		result := removeElement(tt.nums, tt.val)

		if tt.want != result {
			t.Fatalf("want: %d, got: %d", tt.want, result)
		}

		sort.Ints(tt.nums[:result])
		for i := 0; i < result; i++ {
			if tt.nums[i] != tt.wantNums[i] {
				t.Fatalf("wantNums: %+v, got: %+v\n", tt.wantNums, tt.nums)
			}
		}
	}
}

func TestRemoveElement2(t *testing.T) {
	for _, tt := range test {
		result := removeElement2(tt.nums, tt.val)

		if tt.want != result {
			t.Fatalf("want: %d, got: %d", tt.want, result)
		}

		sort.Ints(tt.nums[:result])
		for i := 0; i < result; i++ {
			if tt.nums[i] != tt.wantNums[i] {
				t.Fatalf("wantNums: %+v, got: %+v\n", tt.wantNums, tt.nums)
			}
		}
	}
}
