package plusOne

func plusOne(digits []int) []int {

	var numToAdd int
	digitsLen := len(digits)
	if digits[digitsLen-1]+1 < 10 {
		digits[digitsLen-1] = digits[digitsLen-1] + 1
		return digits
	}

	needTransition := true

	for i := len(digits) - 1; i >= 0; i-- {
		if needTransition {
			numToAdd = digits[i] + 1
			if numToAdd >= 10 {
				numToAdd = numToAdd % 10
				needTransition = true
			} else {
				needTransition = false
			}
		} else {
			return digits
		}
		digits[i] = numToAdd
	}

	if needTransition {
		digits = append([]int{1}, digits...)
	}

	return digits
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func plusOne2(digits []int) []int {
	// reverse
	digitsLen := len(digits)
	if digits[digitsLen-1]+1 < 10 {
		digits[digitsLen-1] = digits[digitsLen-1] + 1
		return digits
	}

	reverse(digits)

	one, i := 1, 0

	for one == 1 {
		if i < len(digits) {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i] += 1
				one = 0
			}
		} else {
			digits = append(digits, 1)
			break
		}
		i++
	}

	reverse(digits)

	return digits

}
