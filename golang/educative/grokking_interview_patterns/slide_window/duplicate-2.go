package main

import (
	"fmt"
	"strings"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if seen[nums[i]] {
			return true
		}

		seen[nums[i]] = true

		if len(seen) > k {
			delete(seen, nums[i-k])
		}
	}

	return false

}

// Driver code
func main() {
	arrs := [][]int{
		{7, 8, 6, 7, 9},
		{-1, 2, -3, 4, -5},
		{900},
		{9, -6, 3, 0, -3, 6, 9},
		{-1000, 1000},
	}
	ks := []int{3, 5, 1, 6, 10000}

	for i, arr := range arrs {
		fmt.Printf("%d.\tarr: [", i+1)
		for j, val := range arr {
			fmt.Print(val)
			if j < len(arr)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
		fmt.Printf("\tk: %d\n", ks[i])
		fmt.Printf("\n\tDo duplicates exist? %s\n", map[bool]string{true: "Yes", false: "No"}[containsNearbyDuplicate(arr, ks[i])])
		fmt.Println(strings.Repeat("-", 100))
	}
}
