package main

import (
	"fmt"
	"strings"
)

func minMovesToMakePalindrome(s string) int {

	chars := []rune(s)
	moves := 0

	for i, j := 0, len(chars)-1; i < j; i++ {
		k := j
		for k > i {
			if chars[i] == chars[k] {
				for k < j {
					chars[k], chars[k+1] = chars[k+1], chars[k]
					moves++
					k++
				}
				j--
				break
			}
			k--
		}
		if k == i {
			moves += len(chars)/2 - i
		}
	}
	return moves

}

func main() {
	strs := []string{"ccxx", "arcacer", "w", "ooooooo", "eggeekgbbeg"}

	for i, s := range strs {
		fmt.Printf("%d.\ts: %s\n", i+1, s)
		fmt.Printf("\tMoves: %d\n", minMovesToMakePalindrome(s))
		fmt.Println(strings.Repeat("-", 100))
	}
}
