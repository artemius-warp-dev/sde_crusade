package main

import (
	"fmt"
	"strings"
)

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	result := make([][]int, 0)

	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastAddedInterval := result[len(result)-1]

		curStart := intervals[i][0]
		curEnd := intervals[i][1]
		prevEnd := lastAddedInterval[1]
		if curStart <= prevEnd {
			if curEnd > prevEnd {
				result[len(result)-1][1] = curEnd
			}
		} else {
			result = append(result, []int{curStart, curEnd})
		}

	}
	return result
}

func main() {
	inputList := [][][]int{
		{
			{1, 5},
			{3, 7},
			{4, 6},
		},
		{
			{1, 5},
			{4, 6},
			{6, 8},
			{11, 15},
		},
		{
			{3, 7},
			{6, 8},
			{10, 12},
			{11, 15},
		},
		{
			{1, 5},
		},
		{
			{1, 9},
			{3, 8},
			{4, 4},
		},
		{
			{1, 2},
			{3, 4},
			{8, 8},
		},
		{
			{1, 5},
			{1, 3},
		},
		{
			{1, 5},
			{6, 9},
		},
		{
			{0, 0},
			{1, 18},
			{1, 3},
		},
	}

	for index, value := range inputList {
		fmt.Printf("%d.\tIntervals to merge: %s\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1))
		result := mergeIntervals(value)
		fmt.Printf("\n\tMerged intervals: [")
		for i, v := range result {
			fmt.Printf("[%d, %d]", v[0], v[1])
			if i != len(result)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("]\n")
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
