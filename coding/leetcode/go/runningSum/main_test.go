package runningSum

import (
	"reflect"
	"testing"
)

var test = []struct {
	nums []int
	want []int
}{
	{
		[]int{1, 2, 3, 4},
		[]int{1, 3, 6, 10},
	},
	{
		[]int{1, 1, 1, 1, 1},
		[]int{1, 2, 3, 4, 5},
	},
	{[]int{3, 1, 2, 10, 1},
		[]int{3, 4, 6, 16, 17},
	},
}

func TestRunningSum(t *testing.T) {

	for _, tt := range test {
		result := runningSum(tt.nums)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Got %v, want %v\n", result, tt.want)
		}
	}
}
