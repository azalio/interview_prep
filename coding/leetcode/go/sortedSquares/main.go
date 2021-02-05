package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	minus := []int{}
	plus := []int{}
	// result := []int{}

	for _, v := range nums {
		if v < 0 {
			minus = append(minus, v*v)
		} else {
			plus = append(plus, v*v)
		}
	}

	// fmt.Println(minus)
	// fmt.Println(plus)

	couter, p, m := 0, 0, len(minus)-1
	for {
		if p >= len(plus) {
			break
		}
		if m < 0 {
			break
		}
		if plus[p] >= minus[m] {
			nums[couter] = minus[m]
			m--
		} else {
			nums[couter] = plus[p]
			p++
		}
		couter++
	}

	for i := p; i < len(plus); i++ {
		nums[couter] = plus[i]
		couter++
	}

	// fmt.Println(m)

	for i := m; i >= 0; i-- {
		// fmt.Println(minus[i])

		nums[couter] = minus[i]
		couter++
	}

	return nums

}

func main() {
	nums := []int{-5, -3, -2, -1}
	result := sortedSquares(nums)
	fmt.Println(result)
}
