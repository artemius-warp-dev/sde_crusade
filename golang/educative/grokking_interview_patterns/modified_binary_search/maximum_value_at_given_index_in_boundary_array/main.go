package main

import "fmt"

func calculateSum(index, mid, n int) int {
	count := 0

	if mid > index {
		count += (mid + (mid - index)) * (index + 1) / 2
	} else {
		count += (mid+1)*mid/2 + (index - mid + 1)
	}

	if mid >= n-index {
		count += (mid + (mid - n + 1 + index)) * (n - index) / 2
	} else {
		count += (mid+1)*mid/2 + (n - index - mid)
	}

	return count - mid

}

func maxValue(n, index, maxSum int) int {
	left, right := 1, maxSum

	for left < right {
		mid := (left + right + 1) / 2

		if calculateSum(index, mid, n) <= maxSum {
			left = mid
		} else {
			right = mid - 1
		}

	}

	return left

}

// Driver code
func main() {
	inputList := [][]int{
		{6, 3, 18},
		{4, 2, 6},
		{3, 0, 3},
		{5, 3, 15},
		{7, 4, 20},
	}

	for i, input := range inputList {
		n := input[0]
		index := input[1]
		maxSum := input[2]
		result := maxValue(n, index, maxSum)
		fmt.Printf("%d.\tInput: n = %d, index = %d, maxSum = %d\n", i+1, n, index, maxSum)
		fmt.Printf("\tMaximum mid at index %d: %d\n", index, result)
		fmt.Println(string('-' * 100))
	}
}
