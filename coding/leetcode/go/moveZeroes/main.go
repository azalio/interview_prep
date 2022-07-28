package moveZeroes

func moveZeroes(nums []int) {
	var l, r int

	for r < len(nums) {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
}
