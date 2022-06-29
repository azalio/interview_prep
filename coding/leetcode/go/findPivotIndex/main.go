package findPivotIndex

func pivotIndex(nums []int) int {

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	// left 0
	if nums[len(nums)-1]-nums[0] == 0 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[len(nums)-1]-nums[i] {
			return i
		}
	}

	return -1
}
