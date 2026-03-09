package main

import "fmt"

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func main() {
	testCases := [][]byte{
		[]byte("hello"),
		[]byte("morning"),
		[]byte("python"),
		[]byte("a"),
		[]byte("racecar"),
	}

	for i, bytes := range testCases {
		original := make([]byte, len(bytes))
		copy(original, bytes)

		reverseString(bytes)

		fmt.Printf("%d\tInput string: [", i+1)
		for j, b := range original {
			fmt.Printf("\"%c\"", b)
			if j != len(original)-1 {
				fmt.Print(",")
			}
		}
		fmt.Println("]\n")

		fmt.Printf("\tReversed string: [")
		for j, b := range bytes {
			fmt.Printf("\"%c\"", b)
			if j != len(bytes)-1 {
				fmt.Print(",")
			}
		}
		fmt.Println("]")

		fmt.Println("----------------------------------------------------------------------------------------------------\n")
	}
}
