package main

import (
	"fmt"
	"strings"
)

func countSubarrays(nums []int, k int64) int64 {
	var result int64 = 0
	var runningSum int64 = 0
	left := 0

	for right := 0; right < len(nums); right++ {
		runningSum += int64(nums[right])

		score := runningSum * int64(right-left+1)

		for left <= right && score >= k {
			runningSum -= int64(nums[left])
			left++

			score = runningSum * int64(right-left+1)
		}

		result += int64(right - left + 1)
		//fmt.Printf("%d, %d : %d %d\n", result, score, right, left)
	}

	return result
}

// Driver code
func main() {
	testArrays := [][]int{
		{2, 1, 4, 3, 5},
		{10, 1, 2},
		{12, 2, 2, 3},
		{5, 4, 2, 10},
		{7, 2, 9, 4, 6},
		{20, 30, 40},
		{11, 1, 3},
		{15, 22, 18, 30, 14, 28, 33, 19, 26, 12},
		{45, 31, 27, 38, 40, 29, 22, 47, 36, 25, 44, 33, 21, 30, 26},
		{100, 200, 300, 400, 500, 99, 98, 97},
	}

	ks := []int64{
		10, 15, 18, 25, 40, 10, 10, 600, 1000, 50,
	}

	for i := 0; i < len(testArrays); i++ {
		nums := testArrays[i]
		k := ks[i]

		fmt.Printf("%d.\tnums: %v\n", i+1, nums)
		fmt.Printf("\tk: %d\n", k)
		fmt.Printf("\n\tCount of subarrays = %d\n", countSubarrays(nums, k))
		fmt.Println(strings.Repeat("-", 100))
	}
}
