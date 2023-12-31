package main

// Idea: two pointer approach
// i (fast pointer) iterates over array, j (slow pointer) shows where to place the next eligible element
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var j int = 2

	for i := 2; i < len(nums); i++ {
		current := nums[i]
		if current == nums[j-2] {
			// third occurence
			continue
		} else if current != nums[j-2] {
			// first or second occurence
			nums[j] = current
			j++
		}
	}

	return j
}
