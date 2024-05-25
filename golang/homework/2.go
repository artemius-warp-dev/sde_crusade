package main

import (
	"fmt"
	"unicode"
)

func printCh(ch rune, times int) {
	for i := 0; i < times; i++ {
		fmt.Printf("%c", ch)
	}

}

func main() {
	test_cases := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
	}

	for _, string := range test_cases {
		stack := []rune{0}
		for _, r := rnge string {
			if r == '\\' {
				stack = append(stack, r)
			} else if unicode.IsDigit(r) && len(stack) > 0 && stack[len(stack)-1] != '\\' {
				printCh(stack[len(stack)-1], int(r-'0')-1)
			} else if unicode.IsDigit(r) && len(stack) > 0 && stack[len(stack)-2] == '\\' {
				printCh(stack[len(stack)-1], int(r-'0'))
			} else if unicode.IsDigit(r) && stack[len(stack)-1] != '\\' {
				fmt.Printf("")
			} else {
				stack = append(stack, r)
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()

	}

	fmt.Println()
}
