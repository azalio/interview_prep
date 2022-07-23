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

func Reverse(input []int) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

func rotate2(nums []int, k int) {

	k = k % len(nums)

	Reverse(nums)

	Reverse(nums[:k])
	Reverse(nums[k:])
}
