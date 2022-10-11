package follow_up_marina

// кстати, советую еще решить похожее - есть массив,
// в котором все числа кроме одного - по 2 раза.
// найти единственное число
// https://leetcode.com/problems/single-number/
// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
// You must implement a solution with a linear runtime complexity and use only constant extra space.

func singleNumber(nums []int) int {
	var addA = make(map[int]int, len(nums)/2+1)

	for _, v := range nums {
		addA[v]++
	}

	for k, v := range addA {
		if v == 1 {
			return k
		}
	}

	return 0
}
