package numberOfGoodPairs

// https://leetcode.com/problems/number-of-good-pairs/
func numIdenticalPairs(nums []int) int {
	// O(n^2)
	ptr := 1
	num := 0

	for _, n := range nums {
		for y := ptr; y < len(nums); y++ {
			if n == nums[y] {
				num++
			}
		}
		ptr++
	}

	return num
}

func numIdenticalPairs2(nums []int) int {
	num := 0
	numMap := make(map[int]int)
	for _, n := range nums {
		num += numMap[n]
		numMap[n]++
	}

	return num
}

func numIdenticalPairs3(nums []int) int {
	num := 0
	numMap := [101]int{}
	for _, n := range nums {
		num += numMap[n]
		numMap[n]++
	}

	return num
}
