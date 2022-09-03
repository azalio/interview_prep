package validPalindrome

import "strings"

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}

func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	s = strip(s)

	y := len(s) - 1

	for x := 0; x < len(s)/2; x++ {
		if s[x] != s[y] {
			return false
		} else {
			y--
		}
	}

	return true
}
