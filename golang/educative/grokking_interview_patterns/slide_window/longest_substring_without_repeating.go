package main

import (
	"fmt"
	"strings"
)

func findLongestSubstring(str string) int {
	if len(str) == 0 {
		return 0
	}

	n := len(str)
	windowStart, longest, windowLength, i := 0, 0, 0, 0
	lastSeenAt := make(map[byte]int)

	for i = 0; i < n; i++ {
		if _, ok := lastSeenAt[str[i]]; !ok {
			lastSeenAt[str[i]] = i
		} else {
			if lastSeenAt[str[i]] >= windowStart {
				windowLength = i - windowStart
				if longest < windowLength {
					longest = windowLength
				}

				windowStart = lastSeenAt[str[i]] + 1
			}
			lastSeenAt[str[i]] = i
		}
	}
	//fmt.Printf("%v\n", lastSeenAt)
	if longest < i-windowStart {
		longest = i - windowStart
	}
	return longest

}

func main() {
	str := []string{"abcabcbb", "pwwkew", "bbbbb", "ababababa", "", "ABCDEFGHI", "ABCDEDCBA", "AAAABBBBCCCCDDDD"}
	for i, s := range str {
		fmt.Printf("%d.\tInput string: \"%s\"\n", i+1, s)
		result := findLongestSubstring(s)
		fmt.Printf("\tLength of the longest substring: %d\n", result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
