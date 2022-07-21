package removeDuplicatesFromArray

func removeDuplicates(nums []int) int {

	var pos1 int

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[pos1] {
			continue
		} else {
			pos1++
			nums[pos1] = nums[i+1]
		}
	}

	return pos1 + 1
}
