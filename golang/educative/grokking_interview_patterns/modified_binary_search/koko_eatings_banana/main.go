package main

import (
	"fmt"
	"strings"
)

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, max(piles)

	for left < right {
		mid := (left + right) / 2
		hours := 0

		for _, pile := range piles {
			hours += (pile + mid - 1) / mid
		}

		if hours <= h {
			right = mid
		} else {
			left = mid + 1
		}

	}
	return left
}

func max(arr []int) int {
	m := arr[0]
	for _, v := range arr {
		if v > m {
			m = v
		}
	}
	return m
}

// Driver code
func main() {
	testCases := []struct {
		piles []int
		h     int
	}{
		{[]int{2, 2, 2, 2, 2}, 5},
		{[]int{9, 18, 27}, 9},
		{[]int{15, 3, 20, 7, 11}, 7},
		{[]int{100, 200, 300, 400}, 20},
		{[]int{6, 6, 6, 6, 6, 6}, 12},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tpiles: %v\n\th: %d\n", i+1, tc.piles, tc.h)
		fmt.Printf("\n\tMinimum eating speed: %d\n", minEatingSpeed(tc.piles, tc.h))
		fmt.Println(strings.Repeat("-", 100))
	}
}
