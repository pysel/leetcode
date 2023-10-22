package easy

// 100% 13%
func removeElement_mine(nums []int, val int) int {
	nlen := len(nums)

	for i := nlen - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i] = nums[nlen-1]
			nums[nlen-1] = val
			nlen--
		}
	}
	return nlen
}
