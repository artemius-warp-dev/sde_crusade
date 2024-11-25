package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	runes := []rune(str)
	count := 0
	left, right := 0, len(runes)-1
	for left < right {
		if count > 1 {
			return false
		}
		if runes[left] == runes[right] {
			left++
			right--
		} else {
			if runes[left] == runes[right-1] {
				right--
			} else if runes[left+1] == runes[right] {
				left++
			}
			count++
		}

	}
	return true
}

func main() {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},  // Standard palindrome.
		{"madam", true},    // Standard palindrome.
		{"level", true},    // Standard palindrome.
		{"", true},         // Empty string is a palindrome.
		{"a", true},        // Single character is a palindrome.
		{"ab", true},       // Removing one character can make it a palindrome.
		{"abc", false},     // Cannot be a palindrome by removing just one character.
		{"abca", true},     // Removing 'c' makes it a palindrome.
		{"deified", true},  // Standard palindrome.
		{"raceecar", true}, // Removing one character makes it a palindrome.
		{"abccba", true},   // Standard palindrome.
		{"abcdef", false},  // Cannot be a palindrome by removing one character.
		{"abcba", true},    // Standard palindrome.
		{"abcdba", true},   // Removing 'd' makes it a palindrome.
	}

	for _, testCase := range testCases {
		result := isPalindrome(testCase.input)
		if result == testCase.expected {
			fmt.Printf("Test passed for input '%s': got %v\n", testCase.input, result)
		} else {
			fmt.Printf("Test failed for input '%s': expected %v, got %v\n", testCase.input, testCase.expected, result)
		}
	}
}
