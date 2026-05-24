package main

import (
	"fmt"
	"math"
	"strings"
)

func splitArray(nums []int, k int) int {
	left, right := math.MinInt, 0
	for _, num := range nums {
		if num > left {
			left = num
		}
		right += num
	}

	for left < right {
		mid := (left + right) / 2

		if canSplit(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canSplit(nums []int, k int, mid int) bool {
	subarrays := 1
	currentSum := 0

	for _, num := range nums {
		if currentSum+num > mid {
			subarrays++
			currentSum = num

			if subarrays > k {
				return false
			}

		} else {
			currentSum += num
		}
	}
	return true
}

// Driver code
func main() {
	splits := [][]int{
		{3, 4, 6, 3},
		{2, 7, 8, 9, 2, 1, 4},
		{12, 53, 43, 67, 35},
		{4, 6, 4, 6, 4, 6},
		{11, 11, 11, 11, 11},
	}
	k := []int{3, 6, 5, 4, 2}

	for i, split := range splits {
		fmt.Printf("%d.\tInput Array: %v\n", i+1, split)
		fmt.Printf("\tk: %d\n", k[i])
		fmt.Printf("\tLargest minimized sum: %d\n", splitArray(split, k[i]))
		fmt.Println(strings.Repeat("-", 100))
	}
}
