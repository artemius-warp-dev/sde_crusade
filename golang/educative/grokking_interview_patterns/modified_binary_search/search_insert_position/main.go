package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	numsList := [][]int{
		{1, 3, 5, 6},
		{1, 3, 5, 6},
		{1, 3, 5, 6},
		{10, 20, 30, 40},
		{2, 4, 6, 8},
	}

	targets := []int{5, 2, 7, 25, 1}

	for i := 0; i < len(numsList); i++ {
		fmt.Println("\tnums: ", numsList[i])
		fmt.Println("\ttarget: ", targets[i])
		result := searchInsert(numsList[i], targets[i])
		fmt.Println("\n\tOutput:", result)
		fmt.Println(string(make([]byte, 100, 100)))
	}
}
