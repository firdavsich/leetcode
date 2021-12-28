package counting_bits

import (
	"fmt"
	"strings"
)

// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n),
// ans[i] is the number of 1's in the binary representation of i.

func countBits(n int) []int {
	var result []int
	for i := 0; i <= n; i++ {
		bin := fmt.Sprintf("%b", i)
		result = append(result, strings.Count(bin, "1"))
	}
	return result
}
