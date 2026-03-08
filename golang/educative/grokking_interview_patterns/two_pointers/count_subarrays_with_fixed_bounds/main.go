package main

import (
	"fmt"
	"strings"
)

func countSubarrays(nums []int, minK int, maxK int) int64 {
	n := len(nums)

	minPos, maxPos, leftBound := -1, -1, -1

	var count int64 = 0

	for i := 0; i < n; i++ {
		if nums[i] < minK || nums[i] > maxK {
			leftBound = i
			minPos, maxPos = -1, -1
		}

		if nums[i] == minK {
			minPos = i
		}

		if nums[i] == maxK {
			maxPos = i
		}

		if minPos != -1 && maxPos != -1 {
			minIndex := min(minPos, maxPos)
			count += int64(max(0, minIndex-leftBound))
		}

	}

	return count

}

// Driver code
func main() {
	testCases := []struct {
		nums []int
		minK int
		maxK int
	}{
		{[]int{1, 3, 5, 2, 7, 5}, 1, 5},
		{[]int{1, 5}, 1, 5},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1, 1},
		{[]int{1, 2, 3, 4}, 2, 5},
		{[]int{1, 5, 1, 5, 1, 5}, 1, 5},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tnums = %v\n\tminK = %d\n\tmaxK = %d\n", i+1, tc.nums, tc.minK, tc.maxK)
		result := countSubarrays(tc.nums, tc.minK, tc.maxK)
		fmt.Printf("\n\tOutput: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
