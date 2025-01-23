package main

import (
	"fmt"
	"strings"
)

func intervalsIntersection(intervalsA [][]int, intervalsB [][]int) [][]int {

	i, j := 0, 0
	intersections := make([][]int, 0)

	for i < len(intervalsA) && j < len(intervalsB) {
		start := max(intervalsA[i][0], intervalsB[j][0])
		end := min(intervalsA[i][1], intervalsB[j][1])

		if start <= end {
			intersections = append(intersections, []int{start, end})
		}

		if intervalsA[i][1] < intervalsB[j][1] {
			i++
		} else {
			j++
		}
	}

	return intersections
}

func max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func main() {
	firstList := [][][]int{
		{
			{1, 2},
		},
		{
			{1, 4},
			{5, 6},
			{9, 15},
		},
		{
			{3, 6},
			{8, 16},
			{17, 25},
		},
		{
			{4, 7},
			{9, 16},
			{17, 28},
			{8, 10},
			{39, 50},
			{55, 66},
			{70, 89},
		},
		{
			{1, 3},
			{5, 6},
			{7, 8},
			{12, 15},
		},
	}

	secondList := [][][]int{
		{
			{1, 2},
		},
		{
			{2, 4},
			{5, 7},
			{9, 15},
		},
		{
			{2, 3},
			{10, 15},
			{18, 23},
		},
		{
			{3, 6},
			{7, 8},
			{9, 10},
			{14, 19},
			{23, 33},
			{35, 40},
			{45, 59},
			{60, 64},
			{68, 76},
		},
		{
			{2, 4},
			{7, 10},
		},
	}

	for i, _ := range firstList {
		intersectingIntervals := intervalsIntersection(firstList[i], secondList[i])
		fmt.Printf("%d.\tInterval A: %s\n", i+1, print(firstList[i]))
		fmt.Printf("\tInterval B: %s\n", print(secondList[i]))
		fmt.Printf("\tIntersecting intervals in 'A' and 'B' are: %s\n", print(intersectingIntervals))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
