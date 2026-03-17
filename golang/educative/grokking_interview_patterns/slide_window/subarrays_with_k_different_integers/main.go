package main

import (
	"fmt"
	"strings"
)

func subarraysWithKDistinct(nums []int, k int) int {
	atMostK := func(nums []int, k int) int {
		count := 0
		left := 0
		freq := make(map[int]int)

		for right := 0; right < len(nums); right++ {
			if freq[nums[right]] == 0 {
				k--
			}
			freq[nums[right]]++

			for k < 0 {
				freq[nums[left]]--
				if freq[nums[left]] == 0 {
					k++
				}
				left++
			}
			count += right - left + 1
		}
		return count
	}

	return atMostK(nums, k) - atMostK(nums, k-1)
}

// Driver code
func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 3, 3}, 2},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{1, 2, 1, 2, 3}, 2},
		{[]int{1, 2, 1, 3, 4}, 3},
	}

	for i, tc := range testCases {
		result := subarraysWithKDistinct(tc.nums, tc.k)
		fmt.Printf("%d.\tnums: %v , k: %d\n", i+1, tc.nums, tc.k)
		fmt.Printf("\tresult: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
