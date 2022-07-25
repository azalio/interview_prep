package plusOne

import (
	"reflect"
	"testing"
)

var test = []struct {
	digits []int
	want   []int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 2, 4},
	},
	{
		[]int{4, 3, 2, 1},
		[]int{4, 3, 2, 2},
	},
	{
		[]int{9},
		[]int{1, 0},
	},
	{
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1},
	},
	{
		[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6},
		[]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7},
	},
	{
		[]int{8, 9, 9, 9},
		[]int{9, 0, 0, 0},
	},
}

func TestPlusOne(t *testing.T) {
	for _, tt := range test {
		oldDigits := make([]int, len(tt.digits))
		copy(oldDigits, tt.digits)
		result := plusOne(tt.digits)
		if !reflect.DeepEqual(tt.want, result) {
			t.Errorf("Want: %v, got: %v\n", tt.want, result)
		}
		copy(tt.digits, oldDigits)
	}
}

func TestPlusOne2(t *testing.T) {
	for _, tt := range test {
		oldDigits := make([]int, len(tt.digits))
		copy(oldDigits, tt.digits)
		result := plusOne2(tt.digits)
		if !reflect.DeepEqual(tt.want, result) {
			t.Errorf("Want: %v, got: %v\n", tt.want, result)
		}
		copy(tt.digits, oldDigits)
	}
}

func BenchmarkPlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			oldDigits := make([]int, len(tt.digits))
			copy(oldDigits, tt.digits)
			result := plusOne(tt.digits)
			if !reflect.DeepEqual(tt.want, result) {
				b.Errorf("Want: %v, got: %v\n", tt.want, result)
			}
			copy(tt.digits, oldDigits)
		}
	}
}

func BenchmarkPlusOne2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			oldDigits := make([]int, len(tt.digits))
			copy(oldDigits, tt.digits)
			result := plusOne2(tt.digits)
			if !reflect.DeepEqual(tt.want, result) {
				b.Errorf("Want: %v, got: %v\n", tt.want, result)
			}
			copy(tt.digits, oldDigits)
		}
	}
}
