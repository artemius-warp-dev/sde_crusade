package main

import (
	"fmt"
	"strings"
)

func findDuplicate(nums []int) int {
	fast, slow := nums[0], nums[0]
	for true {
		slow = nums[slow]
		fast = nums[nums[fast]] // x2 speed
		if slow == fast {
			break
		}
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast] //same speed
	}

	return fast
}

// Driver code
func main() {
	nums := [][]int{
		{1, 3, 2, 3, 5, 4},
		{2, 4, 5, 4, 1, 3},
		{1, 6, 3, 5, 1, 2, 7, 4},
		{1, 2, 2, 4, 3},
		{3, 1, 3, 5, 6, 4, 2},
	}

	for i, num := range nums {
		fmt.Printf("%d.\tnums = %s\n", i+1, strings.Replace(fmt.Sprint(num), " ", ", ", -1))
		fmt.Printf("\tDuplicate number = %d\n", findDuplicate(num))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
