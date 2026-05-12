package main

import (
	"fmt"
	"strings"
)

func numSubarraysWithSum(nums []int, goal int) int {
	start := 0
	prefixZeros := 0
	currentSum := 0
	totalCount := 0

	for end := 0; end < len(nums); end++ {
		currentSum += nums[end]

		for start <= end && currentSum > goal {
			currentSum -= nums[start]
			start++
			prefixZeros = 0
		}

		for start < end && nums[start] == 0 && currentSum == goal {
			prefixZeros++
			currentSum -= nums[start]
			start++
		}

		if start <= end && currentSum == goal {
			totalCount += 1 + prefixZeros
		}

	}

	return totalCount
}

func main() {
	testCases := []struct {
		nums []int
		goal int
	}{
		{[]int{1, 0, 1, 0, 1}, 2},
		{[]int{0, 0, 0, 0, 0}, 0},
		{[]int{1, 1, 1}, 2},
		{[]int{0, 1, 0, 1, 0, 1}, 2},
		{[]int{1}, 1},
	}

	for idx, tc := range testCases {
		result := numSubarraysWithSum(tc.nums, tc.goal)

		fmt.Printf("%d.\tnums: [", idx+1)
		for i, n := range tc.nums {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(n)
		}
		fmt.Println("]")

		fmt.Printf("\tgoal: %d\n", tc.goal)
		fmt.Printf("\tNumber of subarrays with sum = %d are %d.\n", tc.goal, result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
