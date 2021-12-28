package counting_bits

import (
	"fmt"
	"testing"
)

// test countBits function
func TestCountBits(t *testing.T) {
	type test struct {
		input    int
		expected []int
	}
	tests := []test{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, v := range tests {
		testname := fmt.Sprintf("%v", v.input)
		t.Run(testname, func(t *testing.T) {

			actual := countBits(v.input)
			if len(actual) != len(v.expected) {
				t.Errorf("expected %d, actual %d", v.expected, actual)
			}

			for i := 0; i < len(actual); i++ {
				if actual[i] != v.expected[i] {
					t.Errorf("expected %d, actual %d", v.expected, actual)
				}
			}

		})
	}
}
