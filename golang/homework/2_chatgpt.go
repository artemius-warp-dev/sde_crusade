package main

import (
	"fmt"
	"unicode"
)

func printChars(ch rune, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("%c", ch)
	}
}

func main() {
	testCases := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
	}

	for _, str := range testCases {
		stack := []rune{}
		for _, r := range str {
			if r == '\\' {
				// Check if the stack is empty before accessing its last element
				if len(stack) > 0 && stack[len(stack)-1] == '\\' {
					fmt.Printf("\\")
				}
				stack = append(stack, r)
			} else if unicode.IsDigit(r) {
				if len(stack) > 0 && stack[len(stack)-1] == '\\' {
					// Handle escape sequence
					printChars(stack[len(stack)-2], int(r-'0'))
					// Remove the escape character and the digit from the stack
					stack = stack[:len(stack)-2]
				} else {
					// Push the digit onto the stack
					stack = append(stack, r)
				}
			} else {
				// Handle normal characters
				if len(stack) > 0 && stack[len(stack)-1] == '\\' {
					fmt.Printf("%c", r)
					stack = stack[:len(stack)-1] // Remove the escape character from the stack
				} else {
					fmt.Printf("%c", r)
				}
			}
		}
		fmt.Println()
	}
}
