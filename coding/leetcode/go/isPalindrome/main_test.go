package isPalindrome

import "testing"

var test = []struct {
	x    int
	want bool
}{
	{
		121,
		true,
	},
	{
		-121,
		false,
	},
	{
		10,
		false,
	},
	{
		1010101,
		true,
	},
	{
		1111111111011111111,
		false,
	},
	{
		1111111111111111111,
		true,
	},
}

func TestIsPalindrome(t *testing.T) {

	for _, tt := range test {
		result := isPalindrome(tt.x)
		if result != tt.want {
			t.Errorf("x: %d, want: %t, got: %t\n", tt.x, tt.want, result)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {

	for _, tt := range test {
		result := isPalindrome2(tt.x)
		if result != tt.want {
			t.Errorf("x: %d, want: %t, got: %t\n", tt.x, tt.want, result)
		}
	}
}

func BenchmarkTestIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			result := isPalindrome(tt.x)
			if result != tt.want {
				b.Errorf("x: %d, want: %t, got: %t\n", tt.x, tt.want, result)
			}
		}
	}
}

func BenchmarkTestIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			result := isPalindrome2(tt.x)
			if result != tt.want {
				b.Errorf("x: %d, want: %t, got: %t\n", tt.x, tt.want, result)
			}
		}
	}
}
