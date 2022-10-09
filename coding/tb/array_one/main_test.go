package array_one

import (
	"math/rand"
	"testing"
	"time"
)

var resultB int

var test = []struct {
	s    []int
	want int
}{
	{
		[]int{0, 2, 3, 4},
		1,
	},
	{
		[]int{0},
		0,
	},
	//{
	//	[]int{1, 2, 3, 4, 0},
	//	0,
	//},
	{
		[]int{3, 4, 6, 1, 2, 0},
		5,
	},
	//{
	//	[]int{6, 3, 1, 2, 0},
	//	4,
	//},
}

func TestFindNumber(t *testing.T) {
	for _, tt := range test {
		result := findNumber(tt.s)
		if result != tt.want {
			t.Errorf("\nGot %d, want: %d\n", result, tt.want)
		}
	}
}

func TestFindNumber2(t *testing.T) {
	for _, tt := range test {
		result := findNumber2(tt.s)
		if result != tt.want {
			t.Errorf("\nGot %d, want: %d\n", result, tt.want)
		}
	}
}

func TestFindNumber3(t *testing.T) {
	for _, tt := range test {
		result := findNumber3(tt.s)
		if result != tt.want {
			t.Errorf("\nGot %d, want: %d\n", result, tt.want)
		}
	}
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func genBigSlice() []int {
	var result []int
	for i := 0; i < 10000000; i++ {
		result = append(result, i)
	}

	result = remove(result, 31337)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	return result
}

func BenchmarkFindNumber(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		for _, tt := range test {
			r = findNumber(tt.s)
		}

		r = findNumber(genBigSlice())
	}
	resultB = r
}

func BenchmarkFindNumber2(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		for _, tt := range test {
			r = findNumber2(tt.s)
		}
		r = findNumber(genBigSlice())
	}
	resultB = r
}

func BenchmarkFindNumber3(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		for _, tt := range test {
			r = findNumber3(tt.s)
		}
		r = findNumber(genBigSlice())
	}
	resultB = r
}

//go test -bench=. -benchmem
//BenchmarkFindNumber-8                946           1214843 ns/op          813886 B/op        340 allocs/op
//BenchmarkFindNumber2-8               984           1218791 ns/op          814004 B/op        345 allocs/op
