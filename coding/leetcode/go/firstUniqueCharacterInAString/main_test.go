package firstUniqueCharacterInAString

import "testing"

var test = []struct {
	s    string
	want int
}{
	{
		"leetcode",
		0,
	},
	{
		"loveleetcode",
		2,
	},
	{
		"aabb",
		-1,
	},
}

func TestFirstUniqChar(t *testing.T) {
	for _, tt := range test {
		result := firstUniqChar(tt.s)
		if result != tt.want {
			t.Errorf("Want: %d, got: %d\n", tt.want, result)
		}
	}
}
