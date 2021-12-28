package nextGreatestLetter

import (
	"fmt"
	"testing"
)

// test nextGreatestLetter function
func TestNextGreatestLetter(t *testing.T) {
	type test struct {
		input    []byte
		target   byte
		expected byte
	}
	tests := []test{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'a', 'b'}, 'z', 'a'},
	}
	for _, v := range tests {
		testname := fmt.Sprintf("%v", v.input)
		t.Run(testname, func(t *testing.T) {

			actual := nextGreatestLetter(v.input, v.target)
			if actual != v.expected {
				t.Errorf("expected %d, actual %d", v.expected, actual)
			}
		})
	}
}
