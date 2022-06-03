package main

import "testing"

func TestMaxSubArray(t *testing.T) {
	test := []struct {
		nums []int
		want int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{5, 4, -1, 7, 8},
			23,
		},
		{
			[]int{-1},
			-1,
		},
		{
			[]int{-2, -2, -4, -1},
			-1,
		},
	}

	for _, tt := range test {
		result := maxSubArray(tt.nums)
		if result != tt.want {
			t.Errorf("Got: %d, want: %d", result, tt.want)
		}
	}
}
