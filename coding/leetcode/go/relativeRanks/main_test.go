package relativeRanks

import (
	"reflect"
	"testing"
)

var test = []struct {
	score []int
	want  []string
}{
	{
		[]int{5, 4, 3, 2, 1},
		[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
	},
	{
		[]int{10, 3, 8, 9, 4},
		[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
	},
}

func TestFindRelativeRanks(t *testing.T) {
	for _, tt := range test {
		result := findRelativeRanks(tt.score)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("func TestfindRelativeRanks\nwant: %v\ngot:  %v\n",
				tt.want, result)
		}
	}
}
