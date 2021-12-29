package sorted_squares

import (
	"sort"
)

// func sortedSquares(nums []int) []int {
// 	var result []int
// 	for _, v := range nums {
// 		result = append(result, v*v)
// 	}
// 	sort.Ints(result)
// 	return result
// }

func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}
