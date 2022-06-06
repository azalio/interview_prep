package intersection2arr

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) < len(nums2) {
		return intersect(nums2, nums1)
	}

	tempH := make(map[int]int)
	var result []int

	for _, n := range nums2 {
		tempH[n]++
	}

	for _, n := range nums1 {
		v, ok := tempH[n]
		if ok && v != 0 {
			result = append(result, n)
			tempH[n]--
		}
	}
	return result
}

func intersect2(nums1 []int, nums2 []int) []int {
	// Time: O(mlogm + nlogn)
	// Space: O(1)
	sort.Ints(nums1)
	sort.Ints(nums2)

	var result []int

	var i, j int

	for x := 0; x < len(nums1)+len(nums2); x++ {
		if j >= len(nums2) {
			break
		}
		if i >= len(nums1) {
			break
		}
		if nums2[j] == nums1[i] {
			result = append(result, nums2[j])
			j++
			i++
		} else if nums2[j] > nums1[i] {
			i++
		} else {
			j++
		}
	}
	return result
}
