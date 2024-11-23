package main

import (
	"fmt"
	"strings"
)

// Function to check if a number is strobogrammatic
func isStrobogrammatic(num string) bool {
	// Map to store the valid strobogrammatic pairs
	dict := map[rune]rune{
		'0': '0',
		'1': '1',
		'8': '8',
		'6': '9',
		'9': '6',
	}

	left := 0
	right := len(num) - 1

	// Iterate while the left pointer is less than or equal to the right pointer
	for left <= right {
		leftChar := rune(num[left])
		rightChar := rune(num[right])

		// Check if the current digit is valid and matches its corresponding rotated value
		if val, exists := dict[leftChar]; !exists || val != rightChar {
			return false
		}

		// Move pointers towards the center
		left++
		right--
	}

	// Return true if all digit pairs are valid
	return true
}

// Driver code
func main() {
	nums := []string{
		"609",
		"88",
		"962",
		"101",
		"123",
	}

	for i, num := range nums {
		fmt.Printf("%d.\tnum: %s\n", i+1, num)
		fmt.Printf("\n\tIs strobogrammatic: %v\n", isStrobogrammatic(num))
		fmt.Println(strings.Repeat("-", 100))
	}
}
