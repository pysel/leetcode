package main

import "fmt"

// Idea: cyclicly replace values
func rotate_mine(nums []int, k int) {
	var saved_last int = nums[0]
	var saved int
	if len(nums)%2 != 0 {
		for i := 0; i < len(nums); i++ {
			changingInd := (k + i*k) % len(nums)
			saved = nums[changingInd]
			fmt.Println("Index: ", changingInd, "To change: ", nums[changingInd], "Change to: ", saved_last)
			nums[changingInd] = saved_last

			saved_last = saved
		}
		fmt.Println(nums, saved)
	} else {
		for i := 0; i < k; i++ {
			for j := i; j < len(nums)/k; j += k {
				changingInd := (j + j*k) % len(nums)
				saved = nums[changingInd]
				fmt.Println("Index: ", changingInd, "To change: ", nums[changingInd], "Change to: ", saved_last)
				nums[changingInd] = saved_last

				saved_last = saved
			}
		}
	}

	fmt.Println(nums, saved)
}

func rotate(nums []int, k int) {
	nlen := len(nums)
	k %= nlen // nice
	counter := 0
	start := 0

	for counter < nlen {
		changingInd := start
		temp := nums[start]
		for {
			changingInd = (changingInd + k) % nlen
			temp, nums[changingInd] = nums[changingInd], temp
			if start == changingInd {
				break
			}
			counter++
		}
		start++
	}
}

func main() {
	nums := []int{-1, -100, 3, 99}
	k := 2
	rotate(nums, k)
}

// indx:
// init: 1 2 3 4 5 6 7
// outp: 5 6 7 1 2 3 4
