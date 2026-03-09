package main

import (
	"fmt"
	"strings"
)

func appendCharacters(source, target string) int {
	sourceIndex := 0
	targetIndex := 0
	sourceLenght := len(source)
	targetLenght := len(target)

	for sourceIndex < sourceLenght && targetIndex < targetLenght {
		if source[sourceIndex] == target[targetIndex] {
			targetIndex += 1
		}
		sourceIndex += 1
	}

	return targetLenght - targetIndex
}

func main() {
	sources := []string{"axbyc", "abc", "a", "ab", "xyz"}
	targets := []string{"abcde", "abcbc", "a", "aba", "abc"}

	for i := 0; i < len(sources); i++ {
		result := appendCharacters(sources[i], targets[i])
		fmt.Printf("%d\tSource: '%s'\n", i+1, sources[i])
		fmt.Printf("\tTarget: '%s'\n", targets[i])
		fmt.Printf("\tResult: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
