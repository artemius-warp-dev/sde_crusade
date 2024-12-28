package main

import (
	"fmt"
	"strings"
)

func longestRepeatingCharacterReplacement(s string, k int) int {
	stringLength := len(s)
	lengthOfMaxSubstring := 0
	start := 0
	charFreq := make(map[byte]int)
	mostFreqChar := 0

	for end := 0; end < stringLength; end++ {
		if _, ok := charFreq[s[end]]; !ok {
			charFreq[s[end]] = 1
		} else {
			charFreq[s[end]]++
		}

		mostFreqChar = max(mostFreqChar, charFreq[s[end]])

		if end-start+1-mostFreqChar > k {
			charFreq[s[start]]--
			start++
		}

		lengthOfMaxSubstring = max(end-start+1, lengthOfMaxSubstring)
	}

	return lengthOfMaxSubstring
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	inputStrings := []string{"aabccbb", "abbcb", "abccde", "abbcab", "bbbbbbbbb"}
	k := []int{2, 1, 1, 2, 4}

	for i := 0; i < len(inputStrings); i++ {
		fmt.Printf("%d. \t Input String: '%s'\n", i+1, inputStrings[i])
		fmt.Printf("\tk: %d\n", k[i])
		fmt.Printf("\tLenght of the longest substring with repeating characters: %d\n", longestRepeatingCharacterReplacement(inputStrings[i], k[i]))
		fmt.Println(strings.Repeat("-", 100))

	}

}
