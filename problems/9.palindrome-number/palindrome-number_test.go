package palindromenumber

import (
	"testing"
)

// test isPalindrome function
func TestIsPalindrome(t *testing.T) {
	type test struct {
		input    int
		expected bool
	}
	tests := []test{
		{121, true},
		{-121, false},
		{10, false},
		{-101, false},
	}
	for _, v := range tests {
		actual := isPalindrome(v.input)
		if actual != v.expected {
			t.Errorf("isPalindrome(%d): expected %t, actual %t", v.input, v.expected, actual)
		}
	}
}
