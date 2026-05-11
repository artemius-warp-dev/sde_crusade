package main

import (
	"fmt"
	"strings"
)

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	numWords := len(words)
	totalLen := wordLen * numWords

	wordCount := make(map[string]int)

	for _, w := range words {
		wordCount[w]++
	}

	var result []int

	for i := 0; i < wordLen; i++ {
		left := i
		right := i
		seen := make(map[string]int)

		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen

			if _, exists := wordCount[word]; exists {
				seen[word]++

				for seen[word] > wordCount[word] {
					leftWord := s[left : left+wordLen]
					seen[leftWord]--
					left += wordLen
				}

				if right-left == totalLen {
					result = append(result, left)
				}
			} else {
				seen = make(map[string]int)
				left = right
			}
		}
	}
	return result
}

// ===============================
// Driver code to test the function
// ===============================
func main() {
	// 2D array containing lists of words for each test case
	twoDArray := [][]string{
		{"foo", "bar"},                   // Basic example
		{"word", "good", "best", "word"}, // Repeated words
		{"bar", "foo", "the"},            // Multiple valid indices
		{"foo", "bar"},                   // Single match
		{"hi", "jk"},                     // No valid substring
	}

	// Corresponding list of test strings (1D array)
	testStrings := []string{
		"barfoothefoobarman",
		"wordgoodgoodgoodbestword",
		"barfoofoobarthefoobarman",
		"foobar",
		"abcdefg",
	}

	// Loop through both arrays together
	for i := range testStrings {
		result := findSubstring(testStrings[i], twoDArray[i])
		fmt.Println(i + 1)
		fmt.Printf("\tString: %s\n", testStrings[i])
		fmt.Printf("\tWords: %s\n", strings.Join(twoDArray[i], ", "))
		fmt.Printf("\tOutput: %v\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
