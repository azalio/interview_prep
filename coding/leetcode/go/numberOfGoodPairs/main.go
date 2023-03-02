package numberOfGoodPairs

// https://leetcode.com/problems/number-of-good-pairs/
func numIdenticalPairs(nums []int) int {

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
