package singleNumber

func singleNumber(nums []int) int {
	var result int
	for _, n := range nums {
		// we are xor all numbers and get only on who doesn't have a pair
		// 3 xor 3 = 0
		// 001 xor 001 = 000
		result = result ^ n
	}
	return result
}
