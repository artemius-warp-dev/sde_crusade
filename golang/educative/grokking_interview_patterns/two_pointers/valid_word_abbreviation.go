package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func validWordAbbreviation(word string, abbr string) bool {
	rWord := []rune(word)
	rAbbr := []rune(abbr)
	wordPtr, abbrPtr := 0, 0
	strNumber := ""

	for abbrPtr < len(rAbbr) {
		r := rAbbr[abbrPtr]
		fmt.Printf("%c  %d, %d \n", r, abbrPtr, wordPtr)
		if unicode.IsDigit(r) {
			for unicode.IsDigit(rAbbr[abbrPtr]) {
				strNumber += string(rAbbr[abbrPtr])
				abbrPtr++
			}
			fmt.Println(strNumber)
			number, _ := strconv.Atoi(strNumber)
			wordPtr += number

		} else {
			strNumber = ""
			fmt.Printf("CHARACTER: %c, %d, %c\n", r, wordPtr, rWord[wordPtr])
			if r != rWord[wordPtr] {
				return false
			}
			wordPtr++
			abbrPtr++
		}

		if abbrPtr == len(rAbbr)-1 && wordPtr == len(rWord)-1 {
			return true
		}
	}
	fmt.Println("END")
	return false
}

func main() {
	//word, abbr := "physiotherapist", "3sio5pist"
	word, abbr := "internationalization", "13iz4n"
	fmt.Println(validWordAbbreviation(word, abbr))

}
