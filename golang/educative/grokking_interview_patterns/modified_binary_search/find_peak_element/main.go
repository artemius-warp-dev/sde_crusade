package main

import "fmt"

func findPeak(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2

		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	testCases := [][]int{
		{1, 2, 3, 1},          // peak at index 2
		{1, 2, 1, 3, 5, 6, 4}, // peak could be 1 or 5
		{10},                  // single element, peak at 0
		{1, 3, 5, 7, 9},       // strictly increasing, peak at last index
		{9, 7, 5, 3, 1},       // strictly decreasing, peak at first index
	}

	for i, nums := range testCases {
		fmt.Printf("%d.\tnums = %v\n", i+1, nums)
		peakIndex := findPeak(nums)
		fmt.Printf("\tPeak index = %d with Peak value = %d\n", peakIndex, nums[peakIndex])
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
}
