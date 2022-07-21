package rotateArray

func rotate(nums []int, k int) {
	numz := make([]int, len(nums), len(nums))
	var newI int
	_ = copy(numz, nums)
	for i := range numz {
		newI = (i + k) % len(nums)
		nums[newI] = numz[i]
	}
}

//func rotate2(nums []int, k int) {
//	if k >= len(nums) {
//		k = k % len(nums)
//	}
//
//	nums[k:], nums[:k] = nums[:k], nums[k:]
//
//}
