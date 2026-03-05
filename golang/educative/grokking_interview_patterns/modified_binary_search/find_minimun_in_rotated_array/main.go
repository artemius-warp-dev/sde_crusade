package main

import (
	"fmt"
	"strings"
)

func findMin(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		pivot := (low + high) / 2

		if nums[pivot] < nums[high] {
			high = pivot
		} else if nums[pivot] > nums[high] {
			low = pivot + 1
		} else {
			high--
		}
	}

	return nums[low]
}

// Driver code
func main() {
	arrs := [][]int{
		{4, 5, 6, 7, 0, 1, 4},
		{8},
		{500, 600, 700, 800, 900, 1000},
		{-5, -3, -2, 0, 2, 3, 5, 7, 11, -11, -7},
		{64, 128, 256, 256, 512, 2, 4, 8, 16, 16, 16, 32, 64},
	}

	for i, nums := range arrs {
		fmt.Printf("%d.\tnums: [", i+1)
		for j, num := range nums {
			fmt.Print(num)
			if j < len(nums)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]\n")
		fmt.Printf("\tMinimum element: %d\n", findMin(nums))
		fmt.Println(strings.Repeat("-", 100))
	}
}
