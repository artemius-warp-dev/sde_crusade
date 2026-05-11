package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0
	var result strings.Builder

	for i < len(word1) && j < len(word2) {
		result.WriteByte(word1[i])
		result.WriteByte(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		result.WriteByte(word1[i])
		i++
	}

	for j < len(word2) {
		result.WriteByte(word2[j])
		j++
	}
	return result.String()
}

func main() {
	type testCase struct {
		w1, w2 string
	}
	testCases := []testCase{
		{"x", "y"},
		{"hello", "world"},
		{"a", "bcdef"},
		{"zyxwv", "m"},
		{"cat", "dogs"},
	}

	for idx, tc := range testCases {
		result := mergeAlternately(tc.w1, tc.w2)
		fmt.Printf("%d.\tInput array: [\"%s\", \"%s\"]\n", idx+1, tc.w1, tc.w2)
		fmt.Printf("\tResult: \"%s\"\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
