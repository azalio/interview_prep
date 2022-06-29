package isomorficStrings

import "testing"

var test = []struct {
	s    string
	t    string
	want bool
}{
	//{
	//	"egg",
	//	"add",
	//	true,
	//},
	//{
	//	"foo",
	//	"bar",
	//	false,
	//},
	//{
	//	"paper",
	//	"title",
	//	true,
	//},
	{
		"badc",
		"baba",
		false,
	},
}

func TestIsIsomorphic(t *testing.T) {
	for _, tt := range test {
		result := isIsomorphic(tt.s, tt.t)
		if result != tt.want {
			t.Errorf("Got: %t, want: %t\n", result, tt.want)
		}
	}
}
