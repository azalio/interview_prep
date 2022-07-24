package intersectionOfTwoArrays2

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	var result []int

	var i, j int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return result
}

func intersect2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect2(nums2, nums1)
	}

	tempH := make(map[int]int)

	for _, n := range nums1 {
		tempH[n]++
	}

	tempHlen := len(tempH)
	var result []int

	for _, n := range nums2 {
		if tempHlen <= 0 {
			break
		}

		k, ok := tempH[n]
		if ok && k != 0 {
			result = append(result, n)
			tempH[n]--
			if tempH[n] == 0 {
				tempHlen--
			}
		}

	}
	return result
}
