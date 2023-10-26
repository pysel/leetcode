package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 0, 6, 1, 5}
	// merge(a1, b1, &nums)
	fmt.Println(hIndex(nums))
}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	hCandidate := 0
	// 0, 1, 2, 3, 4, 5, 5, 6
	for i := 0; i < len(citations); i++ {
		if citations[i] <= len(citations)-i-1 {
			hCandidate = len(citations) - i - 1
		} else {
			break
		}
	}

	return hCandidate
}

func mergeSort(a *[]int) {
	if len(*a) > 1 {
		mid := len(*a) / 2

		left := make([]int, mid)
		right := make([]int, len(*a)-mid)

		copy(left, (*a)[:mid])
		copy(right, (*a)[mid:])

		mergeSort(&left)
		mergeSort(&right)

		merge(left, right, a)
	}
}

// contract: len(into) == len(a) + len(b)
func merge(a []int, b []int, into *[]int) {
	aI := 0
	bI := 0
	intoI := 0
	for aI < len(a) && bI < len(b) {
		if a[aI] < b[bI] {
			(*into)[intoI] = a[aI]
			aI++
		} else {
			(*into)[intoI] = b[bI]
			bI++
		}
		intoI++
	}

	for aI < len(a) {
		(*into)[intoI] = a[aI]
		aI++
		intoI++
	}

	for bI < len(b) {
		(*into)[intoI] = b[bI]
		bI++
		intoI++
	}
}
