package main

func duplicateZeros(arr []int) {
	possible_dups := 0
	length_ := len(arr) - 1

	for left := 0; left < length_+1; left++ {
		if left > length_-possible_dups {
			break
		}

		if arr[left] == 0 {
			if left == length_-possible_dups {
				arr[length_] = 0
				length_--
				break
			}
			possible_dups++
		}
	}
	// fmt.Printf("possible_dups: %d\n", possible_dups)

	// Start backwards from the last element which would be part of new list.
	last := length_ - possible_dups

	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+possible_dups] = 0
			possible_dups--
			arr[i+possible_dups] = 0
		} else {
			arr[i+possible_dups] = arr[i]

		}
	}

	// fmt.Printf("%v\n", arr)

}

func main() {
	// arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	// duplicateZeros(arr) // [1,0,0,2,3,0,0,4]
	// arr := []int{0, 1, 2, 3}
	// duplicateZeros(arr)
	// arr := []int{0, 0, 0, 0, 0, 0, 0}
	// duplicateZeros(arr)
	// arr := []int{8, 4, 5, 0, 0, 0, 0, 7}
	// duplicateZeros(arr)
	// arr := []int{1, 5, 2, 0, 6, 8, 0, 6, 0}
	// duplicateZeros(arr)
	// arr := []int{8, 4, 5, 0, 0, 0, 0, 7} // [8,4,5,0,0,0,0,0]
	// duplicateZeros(arr)
	arr := []int{9, 0, 9, 0, 6, 0, 0, 0, 1, 1, 6, 5, 4, 4, 8, 3, 0, 0, 0, 1, 5, 3, 0, 0, 7, 2, 1, 0, 2, 0, 9, 1, 0, 2, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 7, 9, 0, 8, 7, 0, 9, 7, 1, 0, 2, 0, 0, 0, 0, 9, 0, 0, 0, 0}
	duplicateZeros(arr)
	// [9,0,0,9,0,0,6,0,0,0,0,0,0,1,1,6,5,4,4,8,3,0,0,0,0,0,0,1,5,3,0,0,0,0,7,2,1,0,0,2,0,0,9,1,0,0,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,6,0,0]
}
