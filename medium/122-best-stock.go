package main

func maxProfit(prices []int) int {
	sum := 0
	max_pointer := 0
	min_pointer := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[max_pointer] {
			max_pointer = i
			continue
		} else if prices[i] < prices[max_pointer] {
			sum += prices[max_pointer] - prices[min_pointer]
			min_pointer, max_pointer = i, i
		}
	}

	sum += prices[max_pointer] - prices[min_pointer]

	return sum
}

// Idea: basic greedy algorithm
func maxProfit_100(prices []int) int {
	buyPrice := prices[0]
	maximumProfit := 0
	for day := 1; day < len(prices); day++ {
		if buyPrice > prices[day] {
			buyPrice = prices[day]
		}
		profit := prices[day] - buyPrice
		if profit > 0 {
			maximumProfit += profit
			buyPrice = prices[day]
		}
	}
	return maximumProfit
}
