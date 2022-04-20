package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	if m == 0 {
		copy(nums1, nums2)
		return
	} else if len(nums2) == 0 {
		return
	}

	pointer1, pointer2 := 0, 0
	new := []int{}
	for {
		// fmt.Printf("pointer1: %d\n", pointer1)
		// fmt.Printf("pointer2: %d\n", pointer2)
		if pointer1 == m {
			for _, v := range nums2[pointer2:] {
				new = append(new, v)
			}
			break
		}
		if pointer2 == n {
			for _, v := range nums1[pointer1:] {
				new = append(new, v)
			}
			break
		}
		if nums1[pointer1] <= nums2[pointer2] {
			new = append(new, nums1[pointer1])
			pointer1++
		} else {
			new = append(new, nums2[pointer2])
			pointer2++
		}
	}
	copy(nums1, new)

}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}
