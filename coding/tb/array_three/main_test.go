package array_three

import (
	"reflect"
	"testing"
)

var test = []struct {
	a    []int
	b    []int
	c    []int
	want []int
}{
	{
		[]int{1, 5, 10, 20, 40, 80},
		[]int{6, 7, 20, 80, 100},
		[]int{3, 4, 15, 20, 30, 70, 80, 120},
		[]int{20, 80},
	},
	{
		[]int{1, 5, 5},
		[]int{3, 4, 5, 5, 10},
		[]int{5, 5, 10, 20},
		[]int{5, 5},
	},
}

func TestFindCommon(t *testing.T) {
	for _, tt := range test {
		result := findCommon(tt.a, tt.b, tt.c)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("\nWant: %v, got: %v\na: %v\nb: %v\nc: %v\n", tt.want, result, tt.a, tt.b, tt.c)
		}
	}
}
