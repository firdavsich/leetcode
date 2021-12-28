package singlenumber

import (
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)

	if len(nums) == 1 {
		return nums[0]
	}

	for i := 1; i < len(nums); i++ {

		if i == len(nums)-1 {
			return nums[i]
		}

		if i == 1 && nums[i-1] != nums[i] {
			return nums[i-1]
		}

		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}
