package main

import (
	"fmt"
	"strings"
)

func checkInclusion(s1, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}

	var s1Counts [26]int
	var windowCounts [26]int

	for i := 0; i < n1; i++ {
		s1Counts[s1[i]-'a']++
	}

	for i := 0; i < n1; i++ {
		windowCounts[s2[i]-'a']++
	}

	if arraysEqual(s1Counts, windowCounts) {
		return true
	}

	for i := n1; i < n2; i++ {
		windowCounts[s2[i]-'a']++
		windowCounts[s2[i-n1]-'a']--
		if arraysEqual(s1Counts, windowCounts) {
			return true
		}
	}

	return false

}

func arraysEqual(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Driver code
func main() {
	testCases := [][2]string{
		{"abc", "okbacof"},        // true
		{"abc", "okbancof"},       // false
		{"adc", "dcda"},           // true
		{"xyz", "x"},              // false
		{"hello", "ooolleoooleh"}, // false
	}

	for i, tc := range testCases {
		s1, s2 := tc[0], tc[1]
		result := checkInclusion(s1, s2)
		fmt.Printf("%d.\ts1 = %q\n\ts2 = %q\n\n\tOutput: %v\n", i+1, s1, s2, result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
