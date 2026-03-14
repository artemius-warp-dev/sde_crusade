package main

import (
	"fmt"
	"strings"
)

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]

		left += 1
		right -= 1
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

// Driver code
func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{-35, -29, -7, 8, 6}, 3},
		{[]int{1}, 5},
		{[]int{10, 20, 30, 40, 50}, 7},
		{[]int{0, 0, 0, 0}, 10},
		{[]int{1, 2}, 1},
		{[]int{2, 4, 6, 8, 10}, 0},
	}

	for i, tc := range testCases {
		numArr := make([]int, len(tc.nums))
		copy(numArr, tc.nums)

		fmt.Printf("%d.\tInput:\n", i+1)
		fmt.Println("\tnums = [" + strings.Trim(strings.Replace(fmt.Sprint(numArr), " ", ", ", -1), "[]") + "]")
		fmt.Printf("\tk = %d\n", tc.k)

		rotate(numArr, tc.k)
		fmt.Println("\n\tOutput = [" + strings.Trim(strings.Replace(fmt.Sprint(numArr), " ", ", ", -1), "[]") + "]")
		fmt.Println(strings.Repeat("-", 100))
	}
}
