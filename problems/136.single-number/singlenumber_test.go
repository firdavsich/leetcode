package singlenumber

import (
	"fmt"
	"testing"
)

// test singleNumber function
func TestSingleNumber(t *testing.T) {
	type test struct {
		input    []int
		expected int
	}
	tests := []test{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}
	for _, v := range tests {
		testname := fmt.Sprintf("%v", v.input)
		t.Run(testname, func(t *testing.T) {

			actual := singleNumber(v.input)
			if actual != v.expected {
				t.Errorf("expected %d, actual %d", v.expected, actual)
			}
		})
	}
}
