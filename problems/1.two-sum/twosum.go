package twosum

func twoSum(nums []int, target int) []int {
	var res []int
	for idx := range nums {
		for id1 := range nums {
			if id1 != idx {
				if nums[idx]+nums[id1] == target {
					res = []int{id1, idx}
				}
			}
		}
	}

	return res
}
