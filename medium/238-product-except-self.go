package main

// first try: run 87% mem 27%
func productExceptSelf_first(nums []int) []int {
	toRight := make([]int, len(nums))
	toLeft := make([]int, len(nums))

	toRightCounter := 1
	toLeftCounter := 1
	for i := 0; i < len(nums); i++ {
		toRightCounter *= nums[i]
		toLeftCounter *= nums[len(nums)-1-i]
		toRight[i] = toRightCounter
		toLeft[len(nums)-1-i] = toLeftCounter
	}

	// cover edge cases
	nums[0] = toLeft[1]
	nums[len(nums)-1] = toRight[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		nums[i] = toRight[i-1] * toLeft[i+1]
	}

	return nums
}
