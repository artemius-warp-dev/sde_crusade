package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	prevEnd := 0

	for _, interval := range intervals {
		end := interval[1]
		if end > prevEnd {
			count++
			prevEnd = end
		}
	}

	return count

}

// Driver code
func main() {
	testCases := [][][]int{
		{{1, 4}, {3, 6}, {2, 8}},
		{{1, 2}, {1, 4}, {3, 4}},
		{{1, 10}, {2, 9}, {3, 8}, {4, 7}},
		{{1, 3}, {4, 6}, {7, 9}},
		{{1, 5}, {2, 3}, {4, 6}},
	}

	for i, intervals := range testCases {
		fmt.Printf("%d.\tIntervals: [", i+1)
		for j, interval := range intervals {
			fmt.Printf("[%d, %d]", interval[0], interval[1])
			if j != len(intervals)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")

		result := removeCoveredIntervals(intervals)
		fmt.Printf("\tResult: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
