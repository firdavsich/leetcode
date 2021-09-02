package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}), []int{5, 6})
	fmt.Println(findDisappearedNumbers([]int{1, 1}), []int{2})

}

func findDisappearedNumbers(nums []int) []int {
	sort.Ints(nums)

	var uniqList = make([]int, len(nums))
	var res []int

	for i := 0; i < len(nums); i++ {
		uniqList[nums[i]-1] = nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if uniqList[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res

}
