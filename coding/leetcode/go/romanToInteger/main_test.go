package romanToInteger

import (
	"reflect"
	"testing"
)

var test = []struct {
	s    string
	want int
}{
	{
		"III",
		3,
	},
	{
		"LVIII",
		58,
	},
	{"MCMXCIV",
		1994,
	},
}

func TestRomanToInt(t *testing.T) {

	for _, tt := range test {
		result := romanToInt(tt.s)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("Input: %s, got: %d, want: %d\n", tt.s, result, tt.want)
		}
	}
}
