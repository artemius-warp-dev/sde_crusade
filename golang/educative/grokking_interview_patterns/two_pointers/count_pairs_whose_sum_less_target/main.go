package main

import (
	"fmt"
	"sort"
	"strings"
)

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	count := 0
	low, high := 0, len(nums)-1

	for low < high {
		if nums[low]+nums[high] < target {
			count += high - low
			low++
		} else {
			high--
		}
	}
	return count
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{10, 1, 6, 2, 3, 8}, 9},
		{[]int{1, 3, 5, 7}, 8},
		{[]int{1, 2, 3, 6}, 6},
		{[]int{2, 4, 6, 8, 10}, 12},
		{[]int{5, 1, 9, 2}, 10},
	}

	for i, testCase := range testCases {
		fmt.Printf("%d\tnums: %v\n", i+1, testCase.nums)
		fmt.Printf("\ttarget: %d\n", testCase.target)

		result := countPairs(testCase.nums, testCase.target)
		fmt.Printf("\n\tNumber of valid pairs: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
