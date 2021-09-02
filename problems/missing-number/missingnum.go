package main

import "fmt"
import "sort"

// func missingNumber(nums []int) int {
// 	sort.Ints(nums)
// 	if nums[0] != 0 {
// 		return 0
// 	}
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i]-1 != nums[i-1] {
// 			return nums[i] - 1
// 		} else {
// 			if i == len(nums)-1 {
// 				return nums[i] + 1
// 			}
// 		}
// 	}
// 	return 1
// }

func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != nums[i-1] {
			return nums[i] - 1
		}
	}
	return 0
}

func main() {
	if missingNumber([]int{3, 0, 1}) != 2 {
		fmt.Println("FAIL1")
		return
	}
	if missingNumber([]int{0, 1}) != 2 {
		fmt.Println("FAIL2")
		return
	}
	if missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) != 8 {
		fmt.Println("FAIL3")
		return
	}
	if missingNumber([]int{0}) != 1 {
		fmt.Println("FAIL4")
		return
	}
	fmt.Println("OK")
}
