package removeElement

import "fmt"

func removeByIndex(nums []int, index int) []int {
	return append(nums[:index], nums[index+1:]...)
}

// 		[]int{0, 1, 2, 2, 3, 0, 4, 2},
//      []int{0, 1, 3, 0, 4}
func removeElement(nums []int, val int) int {
	var result int

	for i := 0; i < len(nums)-result; {
		fmt.Printf("nums: %+v\n", nums)

		if nums[i] == val {
			removeByIndex(nums, i)
			result++
		}
		if nums[i] != val {
			i++
		}
	}
	//fmt.Printf("result: %d\n", result)
	return len(nums) - result
}

func removeElement2(nums []int, val int) int {
	return 1
}
