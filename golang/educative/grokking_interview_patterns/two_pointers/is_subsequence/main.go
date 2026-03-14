package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	sPointer := 0
	tPointer := 0

	for sPointer < len(s) && tPointer < len(t) {
		if s[sPointer] == t[tPointer] {
			sPointer++
		}

		tPointer++
	}

	return sPointer == len(s)
}

// Driver code
func main() {
	type testCase struct {
		s string
		t string
	}

	testCases := []testCase{
		{"abc", "ahbgdc"},
		{"axc", "ahbgdc"},
		{"", "ahbgdc"},
		{"abc", ""},
		{"ace", "abcde"},
	}

	for i, tc := range testCases {
		result := isSubsequence(tc.s, tc.t)
		fmt.Printf("%d.\ts: \"%s\"\n", i+1, tc.s)
		fmt.Printf("\tt: \"%s\"\n", tc.t)
		fmt.Printf("\n\tOutput: %t\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
