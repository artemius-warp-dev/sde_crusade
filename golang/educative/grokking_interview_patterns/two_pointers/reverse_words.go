package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseWords(sentece string) string {
	reg := regexp.MustCompile(`\s+`)
	sentece = reg.ReplaceAllLiteralString(sentece, " ")
	sentece = strings.TrimSpace(sentece)
	runes := []rune(sentece)
	runes = strRev(runes, 0, len(runes)-1)

	strLen := len(runes) - 1

	for start, end := 0, 0; end <= strLen; end++ {
		if end == strLen || runes[end] == ' ' {
			var endIdx int

			if end == strLen {
				endIdx = end
			} else {
				endIdx = end - 1
			}

			runes = strRev(runes, start, endIdx)
			start = end + 1
		}
	}

	return string(runes)

}

func strRev(runes []rune, startRev, endRev int) []rune {
	for startRev < endRev {
		runes[startRev], runes[endRev] = runes[endRev], runes[startRev]
		startRev++
		endRev--
	}
	return runes
}

// Driver code
func main() {
	stringsToReverse := []string{
		"Hello World",
		"a   string   with   multiple   spaces",
		"Case Sensitive Test 1234",
		"a 1 b 2 c 3 d 4 e 5",
		"     trailing spaces",
		"case test interesting an is this",
	}

	for i, str := range stringsToReverse {
		fmt.Printf("%d.\tOriginal string: '%s'\n", i+1, str)
		fmt.Printf("\tReversed string: '%s'\n", reverseWords(str))
		fmt.Println(strings.Repeat("-", 100))
	}
}
