package isomorficStrings

import "testing"

var test = []struct {
	s    string
	t    string
	want bool
}{
	{
		"egg",
		"add",
		true,
	},
	{
		"paper",
		"title",
		true,
	},
	{
		"paper",
		"title",
		true,
	},
	{
		"badc",
		"baba",
		false,
	},
	{
		"bad",
		"bab",
		false,
	},

	{
		"foo",
		"bar",
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

func TestIsIsomorphic2(t *testing.T) {
	for _, tt := range test {
		result := isIsomorphic2(tt.s, tt.t)
		if result != tt.want {
			t.Errorf("Got: %t, want: %t\n", result, tt.want)
		}
	}
}

func BenchmarkTestIsIsomorphic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			result := isIsomorphic(tt.s, tt.t)
			if result != tt.want {
				b.Errorf("Got: %t, want: %t\n", result, tt.want)
			}
		}
	}
}

func BenchmarkTestIsIsomorphic2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range test {
			result := isIsomorphic2(tt.s, tt.t)
			if result != tt.want {
				b.Errorf("Got: %t, want: %t\n", result, tt.want)
			}
		}
	}
}
