package rotateArray

import (
	"reflect"
	"testing"
)

type testData struct {
	nums []int
	k    int
	want []int
}

var test = []testData{
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
		[]int{5, 6, 7, 1, 2, 3, 4},
	},
	{
		[]int{-1, -100, 3, 99},
		2,
		[]int{3, 99, -1, -100},
	},
	{
		[]int{-1},
		2,
		[]int{-1},
	},
}

func TestRotate(t *testing.T) {
	for _, tt := range test {
		oldNums := make([]int, len(tt.nums))
		copy(oldNums, tt.nums)
		rotate(tt.nums, tt.k)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("Want: %v, got: %v", tt.want, tt.nums)
		}
		copy(tt.nums, oldNums)
	}
}

func TestRotate2(t *testing.T) {
	for _, tt := range test {
		oldNums := make([]int, len(tt.nums))
		copy(oldNums, tt.nums)

		rotate2(tt.nums, tt.k)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("Want: %v, got: %v", tt.want, tt.nums)
		}
		copy(tt.nums, oldNums)

	}
}

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			oldNums := make([]int, len(tt.nums))
			copy(oldNums, tt.nums)
			rotate(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				b.Errorf("Want: %v, got: %v", tt.want, tt.nums)
			}
			copy(tt.nums, oldNums)
		}
	}
}

func BenchmarkRotate2(b *testing.B) {
	for i := 0; i < b.N; i++ {

		for _, tt := range test {
			oldNums := make([]int, len(tt.nums))
			copy(oldNums, tt.nums)

			rotate2(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				b.Errorf("Want: %v, got: %v", tt.want, tt.nums)
			}
			copy(tt.nums, oldNums)

		}
	}
}
