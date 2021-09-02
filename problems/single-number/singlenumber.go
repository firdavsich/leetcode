package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	fmt.Printf("Sorted: %v, ", nums)
	if len(nums) == 1 {
		return nums[0]
	}

	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			return nums[i]
		}

		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i-1]

		}
	}
	return -1
}

func main() {
	fmt.Printf("Result: %d, Expect: %d\n", singleNumber([]int{2, 2, 1}), 1)
	fmt.Printf("Result: %d, Expect: %d\n", singleNumber([]int{4, 1, 2, 1, 2}), 4)
	fmt.Printf("Result: %d, Expect: %d\n", singleNumber([]int{1}), 1)

}
