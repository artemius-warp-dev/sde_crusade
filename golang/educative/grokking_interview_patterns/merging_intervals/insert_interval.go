package main

import (
	"fmt"
	"strconv"
	"strings"
)

func insertInterval(existingIntervals [][]int, newInterval []int) [][]int {
	newStart, newEnd := newInterval[0], newInterval[1]

	i, n := 0, len(existingIntervals)
	output := make([][]int, 0)
	for i < n && newStart > existingIntervals[i][0] {
		output = append(output, existingIntervals[i])
		i++
	}

	if len(output) == 0 || output[len(output)-1][1] < newStart {
		output = append(output, []int{newStart, newEnd})
	} else {
		if output[len(output)-1][1] < newEnd {
			output[len(output)-1][1] = newEnd
		}
	}
	for i < n {
		if output[len(output)-1][1] < existingIntervals[i][0] {
			output = append(output, existingIntervals[i])
		} else {
			if existingIntervals[i][1] > output[len(output)-1][1] {
				output[len(output)-1][1] = existingIntervals[i][1]
			}
		}
		i++
	}
	return output
}

func printIntervals(intervalList [][]int) string {
	var out strings.Builder
	out.WriteString("[")
	for i, v := range intervalList {
		out.WriteString("[" + strconv.Itoa(v[0]) + ", " + strconv.Itoa(v[1]) + "]")
		if i != len(intervalList)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString("]")
	return out.String()
}

func main() {
	newIntervals := [][]int{
		{5, 7},
		{8, 9},
		{10, 12},
		{1, 3},
		{1, 10},
	}

	existingIntervals := [][][]int{
		{
			{1, 2},
			{3, 5},
			{6, 8},
		},
		{
			{1, 3},
			{5, 7},
			{10, 12},
		},
		{
			{8, 10},
			{12, 15},
		},
		{
			{5, 7},
			{8, 9},
		},
		{
			{3, 5},
		},
	}

	for i := 0; i < len(newIntervals); i++ {
		fmt.Printf("%d.\tExisting intervals: %s\n", i+1, printIntervals(existingIntervals[i]))
		fmt.Printf("\tNew interval: [%d, %d]\n", newIntervals[i][0], newIntervals[i][1])
		output := insertInterval(existingIntervals[i], newIntervals[i])
		fmt.Printf("\tUpdated intervals: %s\n", printIntervals(output))
		fmt.Println(strings.Repeat("-", 100))
	}
}
