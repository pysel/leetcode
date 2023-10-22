package main

// Idea: the most basic, bad
func removeDuplicates(nums []int) int {
	for i := 0; ; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
