package twosum

import (
	"fmt"
	"testing"
)

// test nextGreatestLetter function
func TestNextGreatestLetter(t *testing.T) {
	type test struct {
		input    []int
		target   int
		expected []int
	}
	tests := []test{
		{
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}
	for _, v := range tests {
		testname := fmt.Sprintf("%v", v.input)
		t.Run(testname, func(t *testing.T) {
			actual := twoSum(v.input, v.target)

			for i, elem := range actual {
				if elem != v.expected[i] {
					t.Errorf("expected %d, actual %d", v.expected, actual)
				}
			}

		})
	}
}
