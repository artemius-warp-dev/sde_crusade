package main

import (
	"fmt"
	"strings"
)

func findMaxAverage(nums []int, k int) float64 {
	currentSum := 0

	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	maxSum := currentSum

	for i := k; i < len(nums); i++ {
		currentSum += nums[i] - nums[i-k]

		if currentSum > maxSum {
			maxSum = currentSum
		}

	}

	return float64(maxSum) / float64(k)

}

// Driver code
func main() {
	inputData := []struct {
		nums []int
		k    int
	}{
		{[]int{10, 5, 2, -1, 6, 3, -2, -4, 4, 1, -3, -6, -1, -2, -5, -7}, 4},
		{[]int{7, 3, 1, -2, 6, 2, -1, -3, 4, 1, -2, -5, 2, 0, -4, -6}, 4},
		{[]int{12, 9, 5, 2, 8, 6, 4, 1, 7, 5, 3, 0, 4, 2, 0, -3}, 4},
		{[]int{-10, -11, -12, -13, -20, -21, -22, -23, -30, -31, -32, -33, -40, -41, -42, -43}, 4},
		{[]int{5, 3, -2, -3, 4, 2, -3, -4, 3, 1, -4, -5, 2, 0, -5, -6}, 4},
	}

	for i, data := range inputData {
		result := findMaxAverage(data.nums, data.k)
		fmt.Printf("%d.\tInput: nums = {%s}, k = %d\n", i+1, intSliceToString(data.nums), data.k)
		fmt.Printf("\tMaximum Average: %.2f\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}

// Helper function to convert a slice of integers to a space-separated string
func intSliceToString(nums []int) string {
	result := ""
	for _, num := range nums {
		result += fmt.Sprintf("%d ", num)
	}
	return strings.TrimSpace(result)
}
