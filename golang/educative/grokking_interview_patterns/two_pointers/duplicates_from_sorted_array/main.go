package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i += 1
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func main() {
	testCases := [][]int{
		{1, 1, 2, 2, 3},
		{-1, -1, 0, 0, 1, 1, 2},
		{5, 5, 5, 5},
		{1, 2, 3, 4},
		{0},
	}

	for idx, nums := range testCases {
		fmt.Printf("%d.\tnums:%v\n", idx+1, nums)

		// because function modifies in-place
		arr := append([]int(nil), nums...)

		k := removeDuplicates(arr)

		fmt.Printf("\n\tUnique Count (k): %d\n", k)
		fmt.Printf("\tArray After Removing Duplicates: %v\n", arr[:k])
		fmt.Println()
		fmt.Println(strings.Repeat("-", 100))
	}
}
