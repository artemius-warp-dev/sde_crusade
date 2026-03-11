package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	n := len(chars)
	w := 0
	i := 0

	for i < n {
		j := i

		for j < n && chars[j] == chars[i] {
			j++
		}

		count := j - i
		chars[w] = chars[i]
		w++

		if count > 1 {
			for _, d := range strconv.Itoa(count) {
				chars[w] = byte(d)
				w++
			}
		}

		i = j

	}

	return w

}

func main() {
	testCases := [][]byte{
		[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, // multiple small runs
		[]byte{'a'}, // single char
		[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, // 12 b's
		[]byte{'x', 'y', 'z'}, // all distinct
		[]byte{'a', 'b', 'b', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}, // long c-run
	}

	for idx, chars := range testCases {
		fmt.Printf("\n%d.\tInput = %c\n", idx+1, chars)
		result := compress(chars)
		fmt.Printf("\n\tCompressed Length = %d\n", result)
		fmt.Printf("\tCompressed Array  = %c\n", chars[:result])
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
}
