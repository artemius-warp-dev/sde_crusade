package main

import (
	"fmt"
	"math"
	"strings"
)

func minSubArrayLen(target int, nums []int) int {
	windowSize := math.MaxInt64

	start, sum := 0, 0

	for end, _ := range nums {
		sum += nums[end]

		for sum >= target {
			currSubArrSize := (end + 1) - start
			windowSize = min(windowSize, currSubArrSize)
			sum -= nums[start]
			start++
		}
	}
	if windowSize != math.MaxInt64 {
		return windowSize
	}

	return 0
}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

// Driver code
func main() {
	targets := []int{7, 4, 11, 10, 5, 15}
	numsList := [][]int{{2, 3, 1, 2, 4, 3}, {1, 4, 4}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 2, 3, 4}, {1, 2, 1, 3}, {5, 4, 9, 8, 11, 3, 7, 12, 15, 44}}
	for i, target := range targets {
		result := minSubArrayLen(target, numsList[i])
		fmt.Printf("%d.\tInput array: %s\n", i+1, strings.Replace(fmt.Sprint(numsList[i]), " ", ", ", -1))
		fmt.Printf("\tTarget: %d\n", target)
		fmt.Printf("\tMinimum Length of Subarray: %d\n", result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
