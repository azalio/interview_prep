package intersection2arr

import (
	"reflect"
	"sort"
	"testing"
)

var test = []struct {
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
	{
		[]int{0, 5, 8, 7, 2, 9, 7, 5},
		[]int{1, 4, 8, 9},
		[]int{8, 9},
	},
}

func TestIntersect(t *testing.T) {
	for _, tt := range test {
		result := intersect(tt.nums1, tt.nums2)
		sort.Ints(result)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Got: %v, want: %v", result, tt.want)
		}
	}

	for _, tt := range test {
		result := intersect2(tt.nums1, tt.nums2)
		sort.Ints(result)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Got: %v, want: %v", result, tt.want)
		}
	}
}

func BenchmarkIntersect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			result := intersect(tt.nums1, tt.nums2)
			sort.Ints(result)
			if !reflect.DeepEqual(result, tt.want) {
				b.Errorf("Got: %v, want: %v", result, tt.want)
			}
		}
	}
}

func BenchmarkIntersect2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			result := intersect2(tt.nums1, tt.nums2)
			sort.Ints(result)
			if !reflect.DeepEqual(result, tt.want) {
				b.Errorf("Got: %v, want: %v", result, tt.want)
			}
		}
	}
}
