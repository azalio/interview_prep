package twoSum2

func twoSum(nums []int, target int) []int {
	tempH := make(map[int]int)

	for i := range nums {
		need := target - nums[i]
		tN, ok := tempH[need]
		if ok {
			return []int{tN, i}
		} else {
			tempH[nums[i]] = i
		}
	}
	return []int{}
}
