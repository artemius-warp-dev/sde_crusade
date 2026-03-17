package main

import (
	"fmt"
	"strings"
)

func findRepeatedDnaSequences(s string) []string {
	toInt := map[rune]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	encodedSequence := make([]int, len(s))

	for i, c := range s {
		encodedSequence[i] = toInt[c]
	}

	k, n := 10, len(s)
	if n <= k {
		return []string{}
	}

	a, h := 4, 0
	a_k := 1
	seenHashes := make(map[int]bool)
	output := make(map[string]bool)

	for i := 0; i < k; i++ {
		h = h*a + encodedSequence[i]
		a_k *= a
	}

	for start := 1; start <= n-k; start++ {
		h = h*a - encodedSequence[start-1]*a_k + encodedSequence[start+k-1]

		if seenHashes[h] {
			output[s[start:start+k]] = true
		} else {
			seenHashes[h] = true
		}

	}

	result := make([]string, 0, len(output))
	for seq := range output {
		result = append(result, seq)
	}

	return result

}

// Driver code
func main() {
	testCases := []string{
		"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
		"AAAAAAAAAAAAA",
		"ACGTACGTACGTACGTACGTACGTACGTACGT",
		"GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG",
		"GTACGTACGTACGCCCCCCCCGGGGG",
	}

	for i, s := range testCases {
		fmt.Printf("%d.\tInput: \"%s\"\n", i+1, s)
		output := findRepeatedDnaSequences(s)
		fmt.Printf("\n\tOutput: [%s]\n", strings.Join(output, ", "))
		fmt.Println(strings.Repeat("-", 100))
	}
}
