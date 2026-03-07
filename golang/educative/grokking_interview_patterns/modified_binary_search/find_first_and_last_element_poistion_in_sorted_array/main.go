package main

import (
	"fmt"
	"strings"
)

func searchRange(nums []int, target int) []int {
	first := findBound(nums, target, true)
	if first == -1 {
		return []int{-1, -1}
	}

	last := findBound(nums, target, false)
	return []int{first, last}
}

// Helper function to find either the first or last occurrence of the target
func findBound(nums []int, target int, findFirst bool) int {
	left := 0
	right := len(nums) - 1

	// Binary search to find either the first or last occurrence
	for left <= right {
		mid := left + (right-left)/2

		// Found the target
		if nums[mid] == target {
			if findFirst {
				// Check if this is the first occurrence
				if mid == left || nums[mid-1] != target {
					return mid
				}
				right = mid - 1 // Keep searching left side
			} else {
				// Check if this is the last occurrence
				if mid == right || nums[mid+1] != target {
					return mid
				}
				left = mid + 1 // Keep searching right side
			}
		} else if nums[mid] < target {
			left = mid + 1 // Move toward the right
		} else {
			right = mid - 1 // Move toward the left
		}
	}

	// Target not found
	return -1
}

// Driver code
func main() {
	numsArr := [][]int{
		{5, 7, 7, 8, 8, 10},
		{5, 7, 7, 8, 8, 10},
		{1, 3, 5, 6, 9},
		{2, 2, 2, 3, 4},
		{1, 2, 3, 4, 5, 5, 5},
	}

	targetArr := []int{8, 6, 5, 2, 5}

	for i := 0; i < len(numsArr); i++ {
		fmt.Printf("%d.\tnums: [", i+1)
		for j := 0; j < len(numsArr[i]); j++ {
			fmt.Print(numsArr[i][j])
			if j < len(numsArr[i])-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
		fmt.Printf("\ttarget: %d\n\n", targetArr[i])

		result := searchRange(numsArr[i], targetArr[i])
		fmt.Printf("\tFirst and last positions: [%d, %d]\n", result[0], result[1])
		fmt.Println(strings.Repeat("-", 100))
	}
}
