package validParentheses

import "testing"

var test = []struct {
	s    string
	want bool
}{
	{
		"()",
		true,
	},
	{
		"()()",
		true,
	},
	{
		"()[]{}",
		true,
	},
	{
		"(]",
		false,
	},
	{
		"([])",
		true,
	},
	{
		"([[]]})",
		false,
	},
	{
		"(()",
		false,
	},
	{
		"((",
		false,
	},
}

func TestIsValid(t *testing.T) {
	for _, tt := range test {
		result := isValid(tt.s)
		if result != tt.want {
			t.Errorf("func TestIsValid\nwant: %t, got %t\n",
				tt.want, result)
		}
	}
}
