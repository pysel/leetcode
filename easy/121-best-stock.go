package main

func maxProfit(prices []int) int {
	min_pointer := 0
	max_pointer := 1
	candidate := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[min_pointer] {
			min_pointer = i
			max_pointer = i
			continue
		}

		if prices[i] > prices[max_pointer] {
			min_pointer = i
		}

		if prices[max_pointer]-prices[min_pointer] > candidate {
			candidate = prices[max_pointer] - prices[min_pointer]
		}
	}

	return candidate
}
