package implementStrStr

import "testing"

var test = []struct {
	haystack string
	needle   string
	want     int
}{
	{
		"sadbutsad",
		"sad",
		0,
	},
	{
		"leetcode",
		"leeto",
		-1,
	},
}

func TestStrStr(t *testing.T) {
	for _, tt := range test {
		result := strStr(tt.haystack, tt.needle)
		if result != tt.want {
			t.Errorf("\nhaystack: %s, needle: %s\nwant: %d, got: %d\n",
				tt.haystack, tt.needle, tt.want, result)
		}
	}
}
