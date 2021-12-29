package sorted_squares

import (
	"fmt"
	"testing"
)

// test sortedSquares function
func TestSortedSquares(t *testing.T) {
	type test struct {
		input    []int
		expected []int
	}
	tests := []test{
		{
			input:    []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			input:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}
	for _, v := range tests {
		testname := fmt.Sprintf("%v", v.input)
		t.Run(testname, func(t *testing.T) {
			actual := sortedSquares(v.input)

			for i := 0; i < len(actual); i++ {
				if actual[i] != v.expected[i] {
					t.Errorf("expected %d, actual %d", v.expected, actual)
				}
			}
		})
	}
}
