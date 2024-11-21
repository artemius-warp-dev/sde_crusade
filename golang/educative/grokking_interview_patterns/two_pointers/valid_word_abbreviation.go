package main

import (
	"fmt"
	"strings"
)

func validWordAbbreviation(word string, abbr string) bool {
	wordIndex := 0
	abbrIndex := 0

	for abbrIndex < len(abbr) {
		if '0' <= abbr[abbrIndex] && abbr[abbrIndex] <= '9' {
			if abbr[abbrIndex] == '0' {
				return false
			}
			num := 0

			for abbrIndex < len(abbr) && '0' <= abbr[abbrIndex] && abbr[abbrIndex] <= '9' {
				num = num*10 + int(abbr[abbrIndex]-'0')
				abbrIndex++
			}
			wordIndex += num
		} else {
			if wordIndex >= len(word) || word[wordIndex] != abbr[abbrIndex] {
				return false
			}
			wordIndex++
			abbrIndex++
		}
	}

	return wordIndex == len(word) && abbrIndex == len(abbr)
}

func main() {
	words := []string{"a", "a", "abcdefghijklmnopqrst", "abcdefghijklmnopqrst", "word", "internationalization", "localization"}
	abbreviations := []string{"a", "b", "a18t", "a19t", "w0rd", "i18n", "l12n"}

	for i, word := range words {
		fmt.Printf("%d.\t word: '%s'\n", i+1, word)
		fmt.Printf("\t abbr: '%s'\n", abbreviations[i])
		fmt.Printf("\n\t Is '%s' a valid abbreviation for the word '%s'? %v\n", abbreviations[i], word, validWordAbbreviation(word, abbreviations[i]))
		fmt.Println(strings.Repeat("-", 100))
	}
}
