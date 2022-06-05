package intersection2arr

import (
	"reflect"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	test := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			[]int{1, 2, 2, 1},
			[]int{2, 2},
			[]int{2, 2},
		},
		{
			[]int{4, 9, 5},
			[]int{9, 4, 9, 8, 4},
			[]int{4, 9},
		},
	}

	for _, tt := range test {
		result := intersect(tt.nums1, tt.nums2)
		sort.Ints(result)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Got: %v, want: %v", result, tt.want)
		}
	}

}
