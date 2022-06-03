package main

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maxSubArray(nums []int) int {

	maxSum := nums[0]
	curSum := nums[0]

	for _, n := range nums[1:] {
		if curSum < 0 {
			curSum = 0
		}
		curSum += n
		maxSum = Max(curSum, maxSum)
	}

	return maxSum
}
