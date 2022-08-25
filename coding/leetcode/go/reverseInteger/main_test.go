package reverseInteger

import "testing"

var test = []struct {
	x    int
	want int
}{
	{
		123,
		321,
	},
	{
		-123,
		-321,
	},
	{
		120,
		21,
	},
	{
		0,
		0,
	},
	{
		1534236469,
		0,
	},
}

func TestReverse(t *testing.T) {
	for _, tt := range test {
		result := reverse(tt.x)
		if result != tt.want {
			t.Errorf("Want: %d, got: %d\n", tt.want, result)
		}
	}
}
