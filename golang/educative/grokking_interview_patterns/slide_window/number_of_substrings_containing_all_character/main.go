package main

import (
	"fmt"
	"strings"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	count := map[byte]int{'a': 0, 'b': 0, 'c': 0}
	left := 0
	result := 0

	for right := 0; right < n; right++ {
		count[s[right]] += 1

		for count['a'] > 0 && count['b'] > 0 && count['c'] > 0 {
			result += n - right
			count[s[left]] -= 1
			left += 1
		}
	}

	return result

}

func main() {
	// Define test cases
	testCases := []string{"abcabc", "aaacb", "abc", "aabbcc", "abababc"}
	for i, s := range testCases {
		result := numberOfSubstrings(s)

		fmt.Printf("%d.\tInput string: \"%s\"\n", i+1, s)
		fmt.Printf("\tResult: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
