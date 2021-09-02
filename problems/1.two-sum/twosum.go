package main

import "fmt"
import "sort"

func twoSum(nums []int, target int) []int {
	var res []int
	for idx, _ := range nums {
		for id1, _ := range nums {
			if id1 != idx {
				if nums[idx]+nums[id1] == target {
					res = []int{idx, id1}
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{0, 4, 3, 0}, 0))
	fmt.Println(twoSum([]int{-3, 4, 3, 90}, 0))

}

// func twoSum(nums []int, target int) []int {
// 	var res, ids []int
// 	fmt.Printf("DEBUG: nums=%v taget=%d\n", nums, target)
// 	for idx, num := range nums {
// 		// if num <= target {
// 		ids = append(ids, idx)
// 		fmt.Printf("DEBUG: idx=%d num=%d\n", idx, num)
// 		// }
// 	}

// 	for _, idx := range ids {
// 		for _, id1 := range ids {
// 			if id1 != idx {
// 				if nums[idx]+nums[id1] == target {
// 					res = []int{idx, id1}
// 				}
// 			}
// 		}

// 	}

// 	return res
// }
