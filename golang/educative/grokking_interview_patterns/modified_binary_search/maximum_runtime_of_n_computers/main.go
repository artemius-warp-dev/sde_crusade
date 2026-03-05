package main

import (
	"fmt"
	"math"
	"strings"
)

func maxRunTime(batteries []int, n int) int {
	totalPower := 0
	for _, b := range batteries {
		totalPower += b
	}

	left, right := 0, totalPower/n

	for left < right {
		mid := right - (right-left)/2

		usable := 0

		for _, b := range batteries {
			usable += int(math.Min(float64(b), float64(mid)))
		}

		if usable >= mid*n {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left

}

func main() {
	testCases := []struct {
		batteries []int
		n         int
	}{
		{[]int{3, 3, 3}, 2},
		{[]int{5}, 1},
		{[]int{1000, 1000, 1000}, 3},
		{[]int{2, 2, 2, 2, 2}, 3},
		{[]int{1, 1, 1, 1, 10}, 2},
	}

	for i, tc := range testCases {
		result := maxRunTime(tc.batteries, tc.n)
		fmt.Printf("%d.\tn: %d, Batteries: %v\n", i+1, tc.n, tc.batteries)
		fmt.Printf("\tresult: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
