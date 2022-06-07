package bestTimeToByAndSell

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	test := []struct {
		prices []int
		want   int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{1, 2, 4},
			3,
		},
		{
			[]int{2, 1, 4},
			3,
		},
		{
			[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			9,
		},
		{
			[]int{1, 2, 1, 2, 1, 3, 5, 1},
			4,
		},
	}

	for _, tt := range test {
		result := maxProfit(tt.prices)
		if result != tt.want {
			t.Errorf("Got: %d, want: %d", result, tt.want)
		}
	}

}
