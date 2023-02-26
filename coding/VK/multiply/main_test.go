package multiply

import (
	"testing"
)

var test = []struct {
	n    int
	want string
}{
	{
		2,
		"1\t2\t\n2\t4\t\n",
	},
	//{
	//	3,
	//	"",
	//},
}

func TestPrintMultTable(t *testing.T) {
	for _, tt := range test {
		result := printMultTable(tt.n)
		if result != tt.want {
			t.Errorf("func printMultTable\nwant:\n%s\ngot:\n%s\n",
				tt.want, result)
		}
	}
}
