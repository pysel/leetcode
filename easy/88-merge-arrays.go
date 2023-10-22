package main

import "fmt"

func merge_mine(nums1 []int, m int, nums2 []int, n int) {
	nums1copy := make([]int, len(nums1))
	copy(nums1copy, nums1)
	fmt.Println(nums1copy)

	if m == 0 {
		nums1 = nums2
	}

	if n == 0 {
		return
	}

	i1 := 0
	i2 := 0
	for i := 0; i < m+n; i++ {
		if i1 == m {
			nums1[i] = nums2[i2]
			continue
		}

		if i2 == n {
			nums1[i] = nums1copy[i1]
			continue
		}

		if nums1copy[i1] < nums2[i2] {
			nums1[i] = nums1copy[i1]
			i1++
		} else {
			nums1[i] = nums2[i2]
			i2++
		}
	}
	fmt.Println(nums1)
}

// Idea: start from the end of both arrays and append to first one sice it is expanded
func merge_best(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
