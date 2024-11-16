package main

import (
	"fmt"
	"strings"
)

func reverseWords(sentence string) string {

	reversedString := reverseString(sentence)

	runes := []rune(reversedString)

	output := ""

	reversedWord := ""

	//fmt.Println(reversedString)

	start, end := 0, 0

	for end < len(runes) {

		//fmt.Printf("%c\n", runes[end])

		if runes[end] == ' ' {

			//fmt.Println(string(runes[start:end]), len(runes[start:end]), end)

			reversedWord = reverseString(string(runes[start:end]))

			//fmt.Println(reversedWord, len([]rune(reversedWord)))

			fmt.Println(reversedWord, len([]rune(reversedWord)))

			output = concat(output, reversedWord)

			start = start + len([]rune(reversedWord)) + 1

			end = start

			//fmt.Println(start, end)

		}

		//fmt.Println(end)

		end++

	}

	reversedWord = reverseString(string(runes[start:end]))

	//fmt.Println(reversedWord, len([]rune(reversedWord)))

	return concat(output, reversedWord)

}

func concat(output string, str string) string {

	var result string

	if output != "" && str != "" {

		output += " "

	}

	//fmt.Println(str, output)

	if str != " " && str != "" {

		result = output + str

	} else {

		result = output

	}

	return result

}

func reverseString(sentence string) string {

	runes := []rune(sentence)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]

	}

	return strings.TrimSpace(string(runes))

}
