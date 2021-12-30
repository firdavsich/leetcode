package backspace_compare

import (
	"fmt"
	"testing"
)

// test backspaceCompare function
func TestBackspaceCompare(t *testing.T) {
	type test struct {
		s        string
		t        string
		expected bool
	}
	tests := []test{
		{
			s:        "ab#c",
			t:        "ad#c",
			expected: true,
		},
		{
			s:        "ab##",
			t:        "c#d#",
			expected: true,
		},
		{
			s:        "a#c",
			t:        "b",
			expected: false,
		},
		{
			s:        "ab##",
			t:        "#c#d#",
			expected: true,
		},
	}
	for _, v := range tests {
		testname := fmt.Sprintf("%v-%v", v.s, v.t)
		t.Run(testname, func(t *testing.T) {

			actual := backspaceCompare(v.s, v.t)
			if actual != v.expected {
				t.Errorf("expected %v, actual %v", v.expected, actual)
			}
		})
	}
}
