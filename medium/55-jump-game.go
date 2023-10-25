package main

// r = 83%; m = 82%
func canJump(nums []int) bool {
	reachable := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if reachable-i <= nums[i] {
			reachable = i
		}
	}

	if reachable == 0 {
		return true
	}

	return false
}
