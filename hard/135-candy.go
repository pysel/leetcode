package main

// Two-pass approach. Idea:
// 1. set all-ones array
// 2. first pass left to right - compare to left neighbor
// 3. second pass right to left - compare to right neighbor
func candy(ratings []int) int {
	returned := make([]int, len(ratings))
	returned[0] = 1
	for i := 1; i < len(ratings); i++ {
		returned[i] = 1

		if ratings[i-1] < ratings[i] {
			returned[i] = returned[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			returned[i] = max(returned[i], returned[i+1]+1)
		}
	}

	sum := 0
	for _, val := range returned {
		sum += val
	}

	return sum
}
