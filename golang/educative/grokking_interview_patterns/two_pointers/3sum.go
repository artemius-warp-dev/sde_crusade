package main

import (
	"fmt"
	"sort"
)

func findSumOfThree(nums []int, target int) bool {
	sort.Ints(nums)
	fmt.Println(nums)

	for i := 0; i < len(nums); i++ {
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

func main() {
	input := []int{3, 4, 1, 0, 4, 1, 2, 2, 6}
	result := findSumOfThree(input, 15)
	fmt.Println(result)
}
