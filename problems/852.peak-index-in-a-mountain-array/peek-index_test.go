package peakindex

import (
	"fmt"
	"testing"
)

// test peakIndexInMountainArray function
func TestPeakIndexInMountainArray(t *testing.T) {
	type test struct {
		input    []int
		expected int
	}
	tests := []test{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
		{[]int{3, 4, 5, 1}, 2},
		{[]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}, 2},
	}
	for _, v := range tests {
		testname := fmt.Sprintf("%v", v.input)
		t.Run(testname, func(t *testing.T) {

			actual := peakIndexInMountainArray(v.input)
			if actual != v.expected {
				t.Errorf("expected %d, actual %d", v.expected, actual)
			}
		})
	}
}
