package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	chars := []byte(s)
	left := 0
	right := len(chars) - 1

	for left < right {
		for left < right && !vowels[chars[left]] {
			left++
		}

		for left < right && !vowels[chars[right]] {
			right--
		}

		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}

	return string(chars)

}

func main() {
	testCases := []string{
		"photosynthesis",
		"celebrity",
		"artificial",
		"intelligence",
		"Python",
	}

	for i, s := range testCases {
		result := reverseVowels(s)
		fmt.Printf("%d.\tInput string: \"%s\"\n", i+1, s)
		fmt.Printf("\tResult: \"%s\"\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
