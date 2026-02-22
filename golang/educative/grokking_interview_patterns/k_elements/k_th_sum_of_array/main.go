package main

import (
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	sum   int
	index int
}

func kSum(nums []int, k int) int {
	maxSum := 0

	for i, val := range nums {
		if val > 0 {
			maxSum += val
		} else {
			nums[i] = -val
		}
	}

	sort.Ints(nums)

	candidates := []Pair{{0, 0}}

	for i := 0; i < k-1; i++ {
		curr := candidates[0]
		candidates = candidates[1:]

		if curr.index < len(nums) {
			new1 := Pair{curr.sum + nums[curr.index], curr.index + 1}
			insertSorted(&candidates, new1)

			if curr.index > 0 {
				new2 := Pair{curr.sum + nums[curr.index] - nums[curr.index-1], curr.index + 1}
				insertSorted(&candidates, new2)
			}

		}
	}
	return maxSum - candidates[0].sum
}

// minHeap simulation
func insertSorted(arr *[]Pair, val Pair) {
	i := sort.Search(len(*arr), func(i int) bool {
		return (*arr)[i].sum > val.sum
	})
	*arr = append(*arr, Pair{})
	copy((*arr)[i+1:], (*arr)[i:])
	(*arr)[i] = val
}

//	{[]int{1, 2, 3}, 3},
// [] → 0
// [1] → 1
// [2] → 2
// [3] → 3
// [1, 2] → 3
// [2, 3] → 5
// [1, 2, 3] → 6

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 3, 3, 3, 0}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{0, 0, 0, 0, 0}, 10},
		{[]int{1, 5, 1, 1}, 3},
		{[]int{2, 2, 2, 2}, 4},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tInput: nums = %v, k = %d\n", i+1, tc.nums, tc.k)
		result := kSum(tc.nums, tc.k)
		fmt.Printf("\tK-th Largest Subsequence Sum: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
