package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 4, 3, 2, 1}
	fmt.Println(productExceptSelf(nums))
}

func productExceptSelf(nums []int) []int {
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
	fmt.Println(toRight)
	fmt.Println(toLeft)

	// toRightCounter := 1
	// for i := 0; i < len(nums); i++ {
	// 	toRightCounter *= nums[i]
	// 	toRight[i] = toRightCounter
	// }

	// toLeftCounter := 1
	// for i := len(nums) - 1; i >= 0; i-- {
	// 	toLeftCounter *= nums[i]
	// 	toLeft[i] = toLeftCounter
	// }
	// fmt.Println(toRight)
	// fmt.Println(toLeft)
	// cover edge cases
	nums[0] = toLeft[1]
	nums[len(nums)-1] = toRight[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		nums[i] = toRight[i-1] * toLeft[i+1]
	}

	return nums
}
