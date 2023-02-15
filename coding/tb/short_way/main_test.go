package short_way

import "testing"

var test = []struct {
	s    string
	want string
}{
	{
		"/a/b/c/./../../d",
		"/a/d",
	},
	{
		"/a/b/c",
		"/a/b/c",
	},
}

func TestShortWay(t *testing.T) {
	for _, tt := range test {
		result := shortWay(tt.s)
		if result != tt.want {
			t.Errorf("\nWant: %s, got: %s", tt.want, result)
		}
	}
}
