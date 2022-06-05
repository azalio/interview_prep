package intersection2arr

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
