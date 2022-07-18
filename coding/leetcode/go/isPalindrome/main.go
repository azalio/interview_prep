package isPalindrome

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	xS := fmt.Sprint(x)

	i, y := 0, len(xS)-1

	for i <= y {
		if xS[i] != xS[y] {
			return false
		}
		i++
		y--
	}

	return true
}

func isPalindrome2(x int) bool {

	if x < 0 {
		return false
	}

	var revNum int

	tmpD := x

	for {
		lastD := tmpD % 10
		revNum = revNum*10 + lastD
		tmpD = tmpD / 10
		if tmpD == 0 {
			break
		}
	}

	if revNum == x {
		return true
	}

	return false
}
