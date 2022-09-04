package stringToInteger_atoi_

import "testing"

var test = []struct {
	s    string
	want int
}{
	{
		"42",
		42,
	},
	{
		"   -42",
		-42,
	},
	{
		"4193 with words",
		4193,
	},
	{
		"words and 987",
		0,
	},
	{
		"-91283472332",
		-2147483648,
	},
	{
		"+-12",
		0,
	},
	{
		"",
		0,
	},
	{
		"  +  413",
		0,
	},
	{
		"   -42",
		-42,
	},
}

func TestMyAtoi(t *testing.T) {
	for _, tt := range test {
		result := myAtoi(tt.s)
		if result != tt.want {
			t.Errorf("\nWant: %d, got: %d\n", tt.want, result)
		}
	}
}
