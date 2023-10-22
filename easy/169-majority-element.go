package main

// Boyer–Moore majority vote algorithm:
// Idea: for every majority element there should be one non-majority element.
// -> After pairing them, there should be left at least one majority element standing
func majorityElement(nums []int) int {
	var candidate int
	var counter int
	for _, num := range nums {
		if counter == 0 {
			candidate = num
			counter = 1
		} else {
			if num == candidate {
				counter++
			} else {
				counter--
			}
		}
	}

	return candidate
}
