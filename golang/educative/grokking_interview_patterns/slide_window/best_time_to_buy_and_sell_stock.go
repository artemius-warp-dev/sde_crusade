package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {

	buy := 0

	profit := 0

	for sell := 1; sell < len(prices); sell++ {
		currentProfit := prices[sell] - prices[buy]
		if prices[buy] < prices[sell] {
			profit = max(profit, currentProfit)
		} else {
			buy = sell
		}

	}

	return profit
}

func main() {
	// Example test case
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Stock prices: %v\n", prices)

	profit := maxProfit(prices)
	fmt.Printf("Maximum profit: %d\n", profit)
}
