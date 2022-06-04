package mergesortedarray

import "testing"
import "reflect"

func TestMerge(t *testing.T) {
	test := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{1, 3, 0, 0},
			2,
			[]int{2, 4},
			2,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 0, 0},
			2,
			[]int{3, 4},
			2,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{3, 4, 0, 0},
			2,
			[]int{1, 2},
			2,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1},
			1,
			[]int{},
			0,
			[]int{1},
		},
		{
			[]int{0},
			0,
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{2, 0},
			1,
			[]int{1},
			1,
			[]int{1, 2},
		},
	}

	for _, tt := range test {
		result := mergeNaive(tt.nums1, tt.m, tt.nums2, tt.n)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Got: %v, want: %v", result, tt.want)
		}
	}

	for _, tt := range test {
		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		if !reflect.DeepEqual(tt.nums1, tt.want) {
			t.Errorf("Got: %v, want: %v", tt.nums1, tt.want)
		}
	}

}
