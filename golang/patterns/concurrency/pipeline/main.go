package main

import (
	"fmt"
	"strings"
)

func stage1(input <-chan string) <-chan string {

	output := make(chan string)

	go func() {
		defer close(output)
		for s := range input {
			output <- strings.ToUpper(s)
		}
	}()

	return output
}

func stage2(input <-chan string) <-chan string {
	output := make(chan string)
	go func() {

		defer close(output)
		for s := range input {
			output <- reverse(s)
		}

	}()

	return output
}

func reverse(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {

	input := make(chan string)

	go func() {
		defer close(input)
		input <- "hello"
		input <- "world"
	}()

	pipeline := stage2(stage1(input))

	for result := range pipeline {
		fmt.Println(result)
	}

}
