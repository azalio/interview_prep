package twoSum

func twoSum(nums []int, target int) []int {
	var result []int
	h := make(map[int]int)

	for i, n := range nums {
		_, ok := h[target-n]
		if ok {
			return []int{h[target-n], i}
		} else {
			h[n] = i
		}
	}

	return result
}
