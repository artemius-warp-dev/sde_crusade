package main

import (
	"fmt"
	"strings"
)

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}

		swap(nums, i, j)
	}

	reverse(nums, i+1, len(nums)-1)
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, start, end int) {
	for start < end {
		swap(nums, start, end)
		start++
		end--
	}
}

func main() {
	testCases := [][]int{
		{4, 1, 5, 2, 9, 3, 7},
		{8, 2, 6, 4, 7, 5},
		{7, 6, 4, 3, 1},
		{2, 6, 8, 7, 8, 7, 9, 4, 1, 2, 4, 5, 8},
		{1, 2},
	}

	for i, nums := range testCases {
		fmt.Printf("%d.\t Original array: %v\n", i+1, nums)
		nextPermutation(nums)
		fmt.Printf("\t Next permutation: %v\n", nums)
		fmt.Println(strings.Repeat("-", 100))
	}
}
