package _85MaxConsecutiveOnes

//Given a binary array nums, return the maximum number of consecutive 1's in the array.

func findMaxConsecutiveOnes(nums []int) int {
	var result int
	var oldResult int

	for _, v := range nums {
		if v == 1 {
			result++
		} else {
			if oldResult < result {
				oldResult = result
			}
			result = 0
		}
	}

	if oldResult > result {
		result = oldResult
	}

	return result
}
