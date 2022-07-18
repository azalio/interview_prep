package bestTimeToBuyAndSellStok2

import "testing"

var test = []struct {
	prices []int
	want   int
}{
	{
		[]int{7, 1, 5, 3, 6, 4},
		7,
	},
	{
		[]int{1, 2, 3, 4, 5},
		4,
	},
	{
		[]int{7, 6, 4, 3, 1},
		0,
	},
}

func TestMaxProfit(t *testing.T) {

	for _, tt := range test {
		result := maxProfit(tt.prices)
		if result != tt.want {
			t.Errorf("Want: %d, got: %d\n", tt.want, result)
		}
	}

}
