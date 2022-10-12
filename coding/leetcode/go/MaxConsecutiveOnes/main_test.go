package MaxConsecutiveOnes

import "testing"

var test = []struct {
	nums []int
	want int
}{
	{
		[]int{1, 1, 0, 1, 1, 1},
		3,
	},
	{
		[]int{1, 0, 1, 1, 0, 1},
		2,
	},
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	for _, tt := range test {
		result := findMaxConsecutiveOnes(tt.nums)
		if result != tt.want {
			t.Errorf("\nGot: %d, want: %d\nArray is: %v\n", result, tt.want, tt.nums)
		}
	}
}

func TestFindMaxConsecutiveOnes2(t *testing.T) {
	for _, tt := range test {
		result := findMaxConsecutiveOnes2(tt.nums)
		if result != tt.want {
			t.Errorf("\nGot: %d, want: %d\nArray is: %v\n", result, tt.want, tt.nums)
		}
	}
}
