package longestCommonPrefix

import "testing"

var test = []struct {
	strs []string
	want string
}{
	{
		[]string{"flower", "flow", "flight"},
		"fl",
	},
	{
		[]string{"dog", "racecar", "car"},
		"",
	},
	{
		[]string{"ab", "a"},
		"a",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, tt := range test {
		result := longestCommonPrefix(tt.strs)
		if result != tt.want {
			t.Errorf("\nSlice is: %v\nwant: %s, got: %s\n",
				tt.strs, tt.want, result)
		}
	}
}

func TestLongestCommonPrefix2(t *testing.T) {
	for _, tt := range test {
		result := longestCommonPrefix2(tt.strs)
		if result != tt.want {
			t.Errorf("\nSlice is: %v\nwant: %s, got: %s\n",
				tt.strs, tt.want, result)
		}
	}
}
