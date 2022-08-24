package reverseString2

import (
	"reflect"
	"testing"
)

var test = []struct {
	input []byte
	want  []byte
}{
	{
		[]byte{'h', 'e', 'l', 'l', 'o'},
		[]byte{'o', 'l', 'l', 'e', 'h'},
	},
	{
		[]byte{'H', 'a', 'n', 'n', 'a', 'h'},
		[]byte{'h', 'a', 'n', 'n', 'a', 'H'},
	},
}

func TestReverseString(t *testing.T) {
	for _, tt := range test {
		input := tt.input
		reverseString(input)
		if !reflect.DeepEqual(input, tt.want) {
			t.Errorf("Want: %v, got: %v", tt.want, input)
		}
	}
}
