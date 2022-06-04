package mergesortedarray

import "fmt"

func insert(a []int, c int, i int) []int {
	return append(a[:i], append([]int{c}, a[i:]...)...)
}

func mergeNaive(nums1 []int, m int, nums2 []int, n int) []int {
	var temp []int
	var i, j int

	if m == 0 {
		return nums2
	}

	if n == 0 {
		return nums1
	}

	for x := 0; x < m+n; x++ {
		if nums1[i] <= nums2[j] {
			temp = append(temp, nums1[i])
			i++
			if i >= m {
				break
			}
		} else {
			temp = append(temp, nums2[j])
			j++
			if j >= n {
				break
			}
		}
	}

	//fmt.Printf("i: %d, j: %d\n", i, j)

	if i < m {
		temp = append(temp, nums1[i:m]...)
	}

	if j < n {
		temp = append(temp, nums2[j:n]...)
	}

	return temp
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	if n == 0 {
		return
	}

	if m == 0 {
		copy(nums1, nums2)
		return
	}

	y := m + n - 1
	i := m - 1
	j := n - 1
	for ; y > 0; y-- {
		//fmt.Printf("i: %d, j: %d\n", i, j)
		//fmt.Printf("nums1[i]: %d, nums2[j]: %d\n", nums1[i], nums2[j])
		if nums1[i] >= nums2[j] {
			nums1[y] = nums1[i]
			i--
			if i < 0 {
				break
			}
		} else {
			nums1[y] = nums2[j]
			j--
			if j < 0 {
				break
			}
		}
	}

	if j >= 0 {
		for y := y - 1; y >= 0; y-- {
			fmt.Printf("j: %d, y: %d\n", j, y)
			nums1[y] = nums2[j]
			j--
		}
	}
	//fmt.Println("===================================")
}
