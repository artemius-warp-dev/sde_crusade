package main

import (
	"fmt"
	"strings"
)

func singleNonDuplicate(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		mid := l + (r-l)/2

		if mid%2 == 1 {
			mid -= 1
		}

		if nums[mid] == nums[mid+1] {
			l = mid + 2
		} else {
			r = mid
		}
	}
	return nums[l]
}

// Driver code
func main() {
	inputList := [][]int{
		{1, 2, 2, 3, 3, 4, 4},
		{1, 1, 2, 2, 3, 4, 4, 5, 5},
		{1, 1, 2, 3, 3},
		{1, 1, 2},
		{0, 2, 2, 3, 3, 4, 4, 5, 5},
	}

	for index, value := range inputList {
		fmt.Printf("%d.\tThe array = %s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		fmt.Printf("\tSingle element found: %d\n", singleNonDuplicate(value))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
