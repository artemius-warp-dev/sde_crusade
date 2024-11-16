package main

import (
	"fmt"
	"strings"
)

func sortColors(colors []int) []int {

	start := 0
	current := 0
	end := len(colors) - 1
	for current <= end {
		if colors[current] == 0 {
			colors[start], colors[current] = colors[current], colors[start]
			start++
			current++
		} else if colors[current] == 1 {
			current++
		} else {
			colors[current], colors[end] = colors[end], colors[current]
			end--
		}
	}

	return colors
}

// Driver code
func main() {
	inputs := [][]int{
		{0, 1, 0},
		{1, 1, 0, 2},
		{2, 1, 1, 0, 0},
		{2, 2, 2, 0, 1, 0},
		{2, 1, 1, 0, 1, 0, 2},
	}

	// Iterate over the inputs and print the sorted array for each
	for i, input := range inputs {
		fmt.Printf("%d.\tcolors: %v\n", i+1, strings.Replace(fmt.Sprint(input), " ", ", ", -1))

		sortedColors := sortColors(input)

		fmt.Printf("\n\tThe sorted array is: %v\n", strings.Replace(fmt.Sprint(sortedColors), " ", ", ", -1))

		fmt.Println(strings.Repeat("-", 100))
	}
}
