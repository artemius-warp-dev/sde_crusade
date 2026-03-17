package main

import (
	"fmt"
	"sort"
	"strings"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	maxFreq := 0
	windowSum := 0

	for right := 0; right < len(nums); right++ {
		target := nums[right]
		windowSum += target

		for (right-left+1)*target > windowSum+k {
			windowSum -= nums[left]
			left++
		}

		if right-left+1 > maxFreq {
			maxFreq = right - left + 1
		}

	}

	return maxFreq
}

// Driver code
func main() {
	testCases := [][]int{
		{1, 2, 4},
		{1, 4, 8, 13},
		{3, 6, 9},
		{2, 3, 5},
		{1, 1, 2},
		{4, 6, 8, 10},
		{10, 12, 5, 1, 15, 20, 13, 4, 7, 3, 9, 14, 2, 8, 6, 16, 11, 18, 19, 17},
		{5, 5, 5, 5, 6, 7, 8, 9, 10, 10, 10, 10, 10, 11, 12, 13, 14, 15, 16, 17},
	}
	kValues := []int{5, 5, 2, 3, 2, 7, 50, 30}

	for i, nums := range testCases {
		fmt.Printf("%d.\tnums: [", i+1)
		for j, num := range nums {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Print(num)
		}
		fmt.Printf("]")
		fmt.Printf("\n\tk = %d\n", kValues[i])
		fmt.Printf("\n\tOutput = %d\n", maxFrequency(nums, kValues[i]))
		fmt.Println(strings.Repeat("-", 100))
	}
}
