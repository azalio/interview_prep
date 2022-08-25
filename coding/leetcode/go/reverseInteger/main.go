package reverseInteger

import "math"

func reverse(x int) int {
	var result int
	isNegative := false

	if x < 0 {
		isNegative = true
		x = -1 * x

	}

	for x > 0 {
		lastDigit := x % 10
		x = x / 10
		if math.MaxInt32-result*10-lastDigit < 0 {
			return 0
		}
		result = result*10 + lastDigit

	}

	if isNegative {
		result = -1 * result
	}

	return result
}
