package main

import (
	"fmt"
	"sort"
	"strings"
)

func findSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[i] + nums[right]
			if sum == target {
				return true
			} else if sum < target {
				left++
			} else {
				right--
			}
		}

	}
	return false
}

// Driver code
func main() {
	numsLists := [][]int{
		{3, 7, 1, 2, 8, 4, 5},
		{-1, 2, 1, -4, 5, -3},
		{2, 3, 4, 1, 7, 9},
		{1, -1, 0},
		{2, 4, 2, 7, 6, 3, 1},
	}
	tList := []int{10, 7, 20, -1, 8}
	for i, nList := range numsLists {
		fmt.Printf("%d. Input array: %s\n", i+1, strings.Replace(fmt.Sprint(nList), " ", ", ", -1))

		if findSumOfThree(nList, tList[i]) {
			fmt.Printf("   Sum for %d exists\n", tList[i])
		} else {
			fmt.Printf("   Sum for %d does not exist\n", tList[i])
		}
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
