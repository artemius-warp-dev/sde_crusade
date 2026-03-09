package main

import (
	"fmt"
	"math"
)

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		if int(math.Abs(float64(nums[left]))) > int(math.Abs(float64(nums[right]))) {
			res[pos] = nums[left] * nums[left]
			left++
		} else {
			res[pos] = nums[right] * nums[right]
			right--
		}

		pos--
	}

	return res
}

func main() {
	testCases := [][]int{
		{-4, -1, 0, 3, 10},   // mix of negatives and positives
		{-7, -3, 2, 3, 11},   // another mix
		{0, 1, 2, 3, 4},      // all non-negative
		{-5, -4, -3, -2, -1}, // all negative
		{1},                  // single element
	}

	for i, nums := range testCases {
		fmt.Printf("%d.\tnums = %v\n", i+1, nums)
		fmt.Printf("\tOutput = %v\n", sortedSquares(nums))
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
}
