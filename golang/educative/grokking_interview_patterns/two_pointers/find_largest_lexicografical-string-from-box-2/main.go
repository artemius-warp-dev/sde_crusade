package main

import (
	"fmt"
	"strings"
)

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}

	n := len(word)
	i, j := 0, 1
	for j < n {
		k := 0
		for j+k < n && word[i+k] == word[j+k] {
			k++
		}

		if j+k < n && word[i+k] < word[j+k] {
			tempIndex := i
			i = j
			j = max(j+1, tempIndex+k+1)
		} else {
			j = j + k + 1
		}
	}

	end := min(n, i+n-numFriends+1)
	return word[i:end]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Driver code
func main() {
	testCases := []struct {
		word       string
		numFriends int
	}{
		{"dbca", 2},
		{"gggg", 4},
		{"acbd", 2},
		{"zxya", 3},
		{"mnopqr", 3},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tInput: word = '%s', numFriends = %d\n", i+1, tc.word, tc.numFriends)
		result := answerString(tc.word, tc.numFriends)
		fmt.Printf("\tLexicographically Largest String: '%s'\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
