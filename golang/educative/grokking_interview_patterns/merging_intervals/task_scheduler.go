package main

import (
	"fmt"
	"sort"
	"strings"
)

func leastTime(tasks []byte, n int) int {
	frequencies := make(map[byte]int)

	for _, t := range tasks {
		frequencies[t] = frequencies[t] + 1
	}

	sortedFrequencies := make([]struct {
		task byte
		freq int
	}, 0, len(frequencies))
	for task, freq := range frequencies {
		sortedFrequencies = append(sortedFrequencies, struct {
			task byte
			freq int
		}{task, freq})
	}

	sort.Slice(sortedFrequencies, func(i, j int) bool {
		return sortedFrequencies[i].freq < sortedFrequencies[j].freq
	})

	maxFreq := sortedFrequencies[len(sortedFrequencies)-1].freq
	sortedFrequencies = sortedFrequencies[:len(sortedFrequencies)-1]
	fmt.Printf("%v\n", sortedFrequencies)

	idleTime := (maxFreq - 1) * n

	for len(sortedFrequencies) > 0 && idleTime > 0 {
		idleTime -= min(maxFreq-1, sortedFrequencies[len(sortedFrequencies)-1].freq)
		sortedFrequencies = sortedFrequencies[:len(sortedFrequencies)-1]
	}

	if idleTime < 0 {
		idleTime = 0
	}

	return len(tasks) + idleTime
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Driver code
func main() {
	allTasks := [][]byte{
		{'A', 'A', 'B', 'B'},
		{'A', 'A', 'A', 'B', 'B', 'C', 'C'},
		{'S', 'I', 'V', 'U', 'W', 'D', 'U', 'X'},
		{'M', 'A', 'B', 'M', 'A', 'A', 'Y', 'B', 'M'},
		{'A', 'K', 'X', 'M', 'W', 'D', 'X', 'B', 'D', 'C', 'O', 'Z', 'D', 'E', 'Q'},
	}
	allNs := []int{2, 1, 0, 3, 3}

	for i := 0; i < len(allTasks); i++ {
		fmt.Printf("%d.\tTasks: ", i+1)
		for j := 0; j < len(allTasks[i]); j++ {
			fmt.Printf("%c", allTasks[i][j])
			if j != len(allTasks[i])-1 {
				fmt.Print(", ")
			}
		}
		fmt.Printf("\n\tn: %d\n", allNs[i])

		minTime := leastTime(allTasks[i], allNs[i])
		fmt.Printf("\tMinimum time required to execute the tasks: %d\n", minTime)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
