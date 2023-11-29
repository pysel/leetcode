package main

// Runtime: 93.30% , Memory: 29.16%
func trap(height []int) int {
	maxInds := maxInds(height)
	if height[maxInds[0]] == 0 {
		return 0 // if max is 0, then all values in array must be zero, so no water can be captured
	}

	leftMaxInd := maxInds[0]
	rightMaxInd := maxInds[len(maxInds)-1]
	water := 0

	// 3 scenarios:
	// 1. scanning from left to first max
	// 2. scanning from right to last max
	// 3. scanning from first max to last max

	var curMaxInd = 0
	var displaced = 0

	// First scanning: capture all water until first max
	if leftMaxInd != 0 {
		for i := 1; i <= leftMaxInd; i++ {
			if height[i] > height[curMaxInd] {
				// Formula: rectArea - displaced
				// rectArea = bottomLength * height
				// height = height[curMaxInd]
				// bottomLength = i - curMaxInd - 1
				water += (i-curMaxInd-1)*height[curMaxInd] - displaced
				displaced = 0
				curMaxInd = i
				continue
			}
			displaced += height[i]
		}
	}

	// Note: variable "displaced" is guaranteed to be zero here

	// Second scanning: capture all water from end to last max
	if rightMaxInd != len(height)-1 {
		curMaxInd = len(height) - 1
		for i := len(height) - 2; i >= rightMaxInd; i-- {
			if height[i] > height[curMaxInd] {
				// Similar to first case, but in the other direction
				water += (curMaxInd-i-1)*height[curMaxInd] - displaced
				displaced = 0
				curMaxInd = i
				continue
			}
			displaced += height[i]
		}
	}

	// Last scanning: from leftMax to RightMax
	if leftMaxInd != rightMaxInd {
		for i := leftMaxInd + 1; i < rightMaxInd; i++ {
			displaced += height[i]
		}

		// Calculate how much water is between edge maxes
		water += (rightMaxInd-leftMaxInd-1)*height[rightMaxInd] - displaced
	}

	return water
}

// maxInds returns indexes of max values
func maxInds(a []int) []int {
	maxInd := 0
	returned := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] > a[maxInd] {
			maxInd = i
			returned = []int{i}
		} else if a[i] == a[maxInd] {
			returned = append(returned, i)
		}
	}

	return returned
}
