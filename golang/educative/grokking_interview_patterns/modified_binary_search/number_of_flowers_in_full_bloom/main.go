package main

import (
	"fmt"
	"sort"
)

func fullBloomFlowers(flowers [][]int, people []int) []int {
	starts := []int{}
	ends := []int{}

	for _, f := range flowers {
		starts = append(starts, f[0])
		ends = append(ends, f[1]+1)
	}

	sort.Ints(starts)
	sort.Ints(ends)

	ans := make([]int, len(people))
	for i, person := range people {
		started := binarySearch(starts, person)
		ended := binarySearch(ends, person)

		ans[i] = started - ended
	}
	return ans
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if target < arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// Driver code
func main() {
	testCases := []struct {
		flowers [][]int
		people  []int
	}{
		{[][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}, []int{2, 3, 7, 11}},
		{[][]int{{1, 10}, {3, 3}}, []int{3, 3, 2}},
		{[][]int{{5, 5}, {5, 5}, {5, 5}}, []int{5, 4, 6}},
		{[][]int{{2, 4}, {6, 8}}, []int{1, 5, 9}},
		{[][]int{{1, 4}, {2, 5}, {3, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tflowers: %v, people: %v\n", i+1, tc.flowers, tc.people)
		res := fullBloomFlowers(tc.flowers, tc.people)
		fmt.Printf("\n\tFlowers in bloom for each person: %v\n", res)
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
}
