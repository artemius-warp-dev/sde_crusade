package main

import (
	"fmt"
	"strings"
)

func circularArrayLoop(nums []int) bool {
	size := len(nums)

	// Iterate through each index of the array 'nums'.
	for i := 0; i < size; i++ {

		// Set slow and fast pointer at current index value.
		slow, fast := i, i

		// Set true in 'forward' if element is positive, set false otherwise.
		forward := nums[i] > 0

		for {
			// Move slow pointer to one step.
			slow = nextStep(slow, nums[slow], size)
			// If cycle is not possible, break the loop and start from next element.
			if isNotCycle(nums, forward, slow) {
				break
			}

			// First move of fast pointer.
			fast = nextStep(fast, nums[fast], size)
			// If cycle is not possible, break the loop and start from next element.
			if isNotCycle(nums, forward, fast) {
				break
			}

			// Second move of fast pointer.
			fast = nextStep(fast, nums[fast], size)
			// If cycle is not possible, break the loop and start from next element.
			if isNotCycle(nums, forward, fast) {
				break
			}

			// At any point, if fast and slow pointers meet each other,
			// it indicate that loop has been found, return True.
			if slow == fast {
				return true
			}
		}
	}

	return false
}

// A function to calculate the next step
func nextStep(pointer int, value int, size int) int {
	result := (pointer + value) % size
	if result < 0 {
		result += size
	}
	return result
}

// A function to detect a cycle doesn't exist
func isNotCycle(nums []int, prevDirection bool, pointer int) bool {
	// Set current direction to True if current element is positive, set False otherwise.
	currDirection := nums[pointer] >= 0

	// If current direction and previous direction is different or moving a pointer takes back to the same value,
	// then the cycle is not possible, we return True, otherwise return False.
	if (prevDirection != currDirection) || (abs(nums[pointer])%len(nums) == 0) {
		return true
	} else {
		return false
	}
}

// abs function to calculate absolute value of the given integer
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Driver code
func main() {
	inputArray := [][]int{
		{-2, -3, -9},
		{-5, -4, -3, -2, -1},
		{-1, -2, -3, -4, -5},
		{2, 1, -1, -2},
		{-1, -2, -3, -4, -5, 6},
		{1, 2, -3, 3, 4, 7, 1},
		{2, 2, 2, 7, 2, -1, 2, -1, -1}}

	for i, arr := range inputArray {
		fmt.Printf("%d.\tCircular array = %s\n", i+1, strings.Replace(fmt.Sprint(arr), " ", ", ", -1))
		fmt.Printf("\n\tFound loop = %v\n", circularArrayLoop(arr))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
