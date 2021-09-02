package main

import "fmt"
import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	last := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != last {
			last = nums[i]
		} else {
			return true
		}
	}
	return false
}

func main() {
	if containsDuplicate([]int{1, 2, 3, 1}) != true {
		fmt.Println("FAIL1")
		return
	}
	if containsDuplicate([]int{1, 2, 3, 4}) != false {
		fmt.Println("FAIL2")
		return
	}
	if containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) != true {
		fmt.Println("FAIL3")
		return
	}
	fmt.Println("OK")
}
