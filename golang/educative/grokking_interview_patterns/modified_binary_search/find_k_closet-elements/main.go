package main

import (
	"fmt"
	"strings"
)

func binarySearch(array []int, target int) int {
	// Initialize the left and right pointer
	left := 0
	right := len(array) - 1

	// Find the first closest element to target while left is less than or equal to right
	for left <= right {
		// Compute the middle index
		mid := (left + right) / 2

		// If the value at mid is equal to target, return mid
		if array[mid] == target {
			return mid
		}

		// If the value at mid is less than target, move left forward
		if array[mid] < target {
			left = mid + 1
		} else { // If the value at mid is greater than target, move right backward
			right = mid - 1
		}
	}

	// If the target is not found, return the left index (position where the target should be inserted)
	return left
}

func findClosestElements(nums []int, k int, target int) []int {
	if len(nums) == k {
		return nums
	}

	if target <= nums[0] {
		return nums[0:k]
	}

	if target >= nums[len(nums)-1] {
		return nums[len(nums)-k : len(nums)]
	}

	firstClosest := binarySearch(nums, target)

	windowLeft := firstClosest - 1
	windowRight := windowLeft + 1

	for (windowRight - windowLeft - 1) < k {
		if windowLeft == -1 {
			windowRight += 1
			continue
		}

		if windowRight == len(nums) || abs(nums[windowLeft]-target) <= abs(nums[windowRight]-target) {
			windowLeft -= 1
		} else {
			windowRight += 1
		}
	}
	return nums[windowLeft+1 : windowRight]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Driver code
func main() {
	inputList := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5},
		{1, 2, 4, 5, 6},
		{1, 2, 3, 4, 5, 10},
	}
	k := []int{4, 4, 2, 3}
	num := []int{4, 3, 10, -5}

	for index, value := range inputList {
		fmt.Printf("%d.\tThe %d closest elements for the number %d in the array %s are:\n", index+1, k[index], num[index], strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\t %s\n", strings.Replace(fmt.Sprint(findClosestElements(value, k[index], num[index])), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
