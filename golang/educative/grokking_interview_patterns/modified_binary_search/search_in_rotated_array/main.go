package main

import (
	"fmt"
	"strings"
)

func binarySearchRotated(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < target && target <= nums[high] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Driver code
func main() {
	numsList := [][]int{
		{5, 6, 7, 1, 2, 3, 4},
		{40, 50, 60, 10, 20, 30},
		{47, 58, 69, 72, 83, 94, 12, 24, 35},
		{77, 82, 99, 105, 5, 13, 28, 41, 56, 63},
		{48, 52, 57, 62, 68, 72, 5, 7, 12, 17, 21, 28, 33, 37, 41},
	}
	targetList := []int{1, 50, 12, 56, 5}

	for index, value := range numsList {
		fmt.Printf("%d.\tRotated array: %s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\tTarget %d found at index %d\n", targetList[index], binarySearchRotated(value, targetList[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
