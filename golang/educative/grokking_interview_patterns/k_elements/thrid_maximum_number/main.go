package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func thirdMax(nums []int) int {
	heap := []int{}

	taken := make(map[int]bool)

	for _, num := range nums {
		if taken[num] {
			continue
		}

		if len(heap) == 3 {
			if heap[0] < num {
				delete(taken, heap[0])
				heap = heap[1:]

				heap = append(heap, num)
				taken[num] = true
				sort.Ints(heap)
			}
		} else {
			heap = append(heap, num)
			taken[num] = true

			sort.Ints(heap)
		}
	}

	if len(heap) == 1 {
		return heap[0]
	} else if len(heap) == 2 {
		firstNum := heap[0]
		return int(math.Max(float64(firstNum), float64(heap[1])))
	}

	return heap[0]
}

// Driver code
func main() {
	testCases := [][]int{
		{3, 2, 1},
		{1, 2},
		{2, 2, 3, 1},
		{5, 5, 4, 3, 2},
		{1, 1, 1, 1},
	}

	// Iterate over test cases and display results
	for i, testCase := range testCases {
		fmt.Printf("%d. \tnums: %v\n", i+1, testCase)
		fmt.Printf("\tThe third maximum is: %d\n", thirdMax(testCase))
		fmt.Println(strings.Repeat("-", 100))
	}
}
