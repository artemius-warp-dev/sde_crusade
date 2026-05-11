package main

import (
	"fmt"
	"strings"
)

func numberOfSubstrings(s string, k int) int64 {
	n := len(s)
	freq := make([]int, 26)
	charsWithK := 0
	var total int64 = 0
	left := 0

	for right := 0; right < n; right++ {
		charFreqIdx := int(s[right] - 'a')
		freq[charFreqIdx]++
		if freq[charFreqIdx] == k {
			charsWithK++
		}

		for charsWithK > 0 {
			total += int64(n - right)
			leftCharFreqIdx := int(s[left] - 'a')
			if freq[leftCharFreqIdx] == k {
				charsWithK--
			}
			freq[leftCharFreqIdx]--
			left++
		}
	}
	return total
}

// Driver code
func main() {
	testCases := []struct {
		s string
		k int
	}{
		{"abcabc", 2},
		{"vvvvvvvv", 2},
		{"xyxyxyxy", 3},
		{"mnolllonm", 3},
		{"abcdefg", 1},
	}

	for i, tc := range testCases {
		result := numberOfSubstrings(tc.s, tc.k)
		fmt.Printf("%d:\ts = \"%s\"\n\tk = %d\n\n\tOutput = %d\n", i+1, tc.s, tc.k, result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
