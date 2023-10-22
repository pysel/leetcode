package main

// func removeDuplicates(nums []int) int {
// 	var k int
// 	for i := 0; i < len(nums)-2; i++ {
// 		first, second, third := nums[i], nums[i+1], nums[i+2]
// 		if first == second && first == third {
// 			// move last element of the triplet to the end of array
// 			nums = append(nums[:i+2], append(nums[i+3:], nums[i+2])...)
// 			k++
// 			i--
// 		}
// 	}

// 	return k
// }
