package main

// 75%  ||   46%
// Idea: greedy algorithm, for every value, calculate the maximum value we could get to by iterating over
// all values to which we can jump from current index
func jump(nums []int) int {
	jumps := 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] >= len(nums)-1 {
			jumps++
			break
		}

		candidate := i + 1
		for j := i + 1; j < min(i+1+nums[i], len(nums)); j++ {
			if j+nums[j] > candidate+nums[candidate] {
				candidate = j
			}
		}
		jumps++
		if candidate == len(nums)-1 {
			break
		}

		i = candidate - 1
	}

	return jumps
}
