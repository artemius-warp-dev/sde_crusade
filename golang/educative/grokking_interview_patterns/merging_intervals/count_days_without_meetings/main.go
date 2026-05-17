package main

import (
	"fmt"
	"sort"
	"strings"
)

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	occupied := 0

	start, end := meetings[0][0], meetings[0][1]

	for i := 1; i < len(meetings); i++ {
		if meetings[i][0] <= end {
			end = max(end, meetings[i][1])
		} else {
			occupied += (end - start + 1)

			start = meetings[i][0]
			end = meetings[i][1]
		}
	}

	occupied += (end - start + 1)

	return days - occupied
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Driver code
func main() {
	inputDays := []int{12, 6, 100000, 3136, 786}
	inputMeetings := [][][]int{
		{{5, 6}, {9, 11}, {1, 3}},
		{{2, 4}, {5, 5}},
		{{1, 100000}},
		{{361, 570}, {420, 1225}, {72, 144}, {987, 1444}},
		{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}},
	}

	for i, days := range inputDays {
		fmt.Printf("%d.\tdays: %d\n", i+1, days)
		fmt.Printf("\tmeetings: %v\n", inputMeetings[i])
		fmt.Printf("\n\tNumber of free days: %d\n", countDays(days, inputMeetings[i]))
		fmt.Println(strings.Repeat("-", 100))
	}
}
