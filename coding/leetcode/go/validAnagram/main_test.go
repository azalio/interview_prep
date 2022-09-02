package validAnagram

import "testing"

var test = []struct {
	s    string
	t    string
	want bool
}{
	{
		"anagram",
		"nagaram",
		true,
	}, {
		"rat",
		"car",
		false,
	},
	{
		"ab",
		"a",
		false,
	},
}

func TestIsAnagram(t *testing.T) {
	for _, tt := range test {
		result := isAnagram(tt.s, tt.t)
		if result != tt.want {
			t.Errorf("\ns: %s\nt: %s\nwant: %t, got: %t\n", tt.s, tt.t, tt.want, result)
		}
	}
}
