package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findNextPalindrome(numStr string) string {
	runes := []rune(numStr)
	var mid int
	var left, right []rune
	result := ""
	if len(runes)%2 == 0 {
		//even
		left = runes[0 : len(runes)/2]
		right = runes[len(runes)/2:]
		mid = len(runes) / 2

		fmt.Printf("%q, %q, %d\n", left, right, mid)
		for i := 1; i < len(left); i++ {
			if int(left[i]) <= int(left[i-1]) {
				//decreasing digit
				left[i], left[i-1] = left[i-1], left[i]
				// reverse remaining left half
				for j, k := i+1, len(left)-1; j < k; j, k = j+1, k+1 {
					left[j], left[k] = left[k], left[j]
				}
				// mirror palindrom from left to the right
				for i := 0; i < mid-1; i++ {
					fmt.Println(len(left) - 1 - i)
					right[i] = left[len(left)-1-i]
				}
				result = string(left) + string(right)

			}
		}

	} else {
		//odd
		left = runes[0 : len(runes)/2]
		right = runes[len(runes)/2+1:]
		mid = len(runes)/2 + 1

		fmt.Printf("%q, %q, %d\n", left, right, mid)
		for i := 1; i < len(left); i++ {
			if int(left[i]) <= int(left[i-1]) {
				//decreasing digit
				left[i], left[i-1] = left[i-1], left[i]
				// reverse remaining left half
				for j, k := i+1, len(left)-1; j < k; j, k = j+1, k+1 {
					left[j], left[k] = left[k], left[j]
				}
				// mirror palindrom from left to the right
				for i := 0; i < mid-1; i++ {
					fmt.Println(len(left) - 1 - i)
					right[i] = left[len(left)-1-i]
				}
				result = string(left) + string(runes[mid]) + string(right)
			}
		}

	}

	//fmt.Printf("After permutation: %q, %q\n", left, right)

	resultInt, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println("Error converting result to integer:", err)
		return ""
	}

	inputInt, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting result to integer:", err)
		return ""
	}

	if resultInt < inputInt {
		return result
	}

	// Replace the following return statement with your code
	return ""
}

func main() {
	testCases := []string{"1221", "54345", "999", "12321", "89798"}

	for i, testCase := range testCases {
		fmt.Printf("%d.\t Original palindrome: '%s'\n", i+1, testCase)
		nextPalindrome := findNextPalindrome(testCase)
		fmt.Printf("\t Next greater palindrome: '%s'\n", nextPalindrome)
		fmt.Println(strings.Repeat("-", 100))
	}
}
