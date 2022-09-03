package validPalindrome

import (
	"testing"
)

var test = []struct {
	s    string
	want bool
}{
	{
		"A man, a plan, a canal: Panama",
		true,
	},
	{
		"race a car",
		false,
	},
	{
		" ",
		true,
	},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range test {
		result := isPalindrome(tt.s)
		if result != tt.want {
			t.Errorf("\ns: %s\nwant: %t, got: %t\n", tt.s, tt.want, result)
		}
	}
}
