package main

import (
	"fmt"
	"strings"
)

func findNextPermutation(digits []byte) bool {
	i := len(digits) - 2
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}
	if i == -1 {
		return false
	}

	j := len(digits) - 1
	for digits[j] <= digits[i] {
		j--
	}

	digits[i], digits[j] = digits[j], digits[i]
	reverse(digits[i+1:])
	return true
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func reverseBytes(arr []byte) []byte {
	reversed := make([]byte, len(arr))
	copy(reversed, arr)
	reverse(reversed)
	return reversed
}

func findNextPalindrome(numStr string) string {
	n := len(numStr)

	if n == 1 {
		return ""
	}

	halfLength := n / 2
	leftHalf := []byte(numStr[:halfLength])

	if !findNextPermutation(leftHalf) {
		return ""
	}

	var nextPalindrome string
	if n%2 == 0 {
		nextPalindrome = string(leftHalf) + string(reverseBytes(leftHalf))
	} else {
		middleDigit := numStr[halfLength]
		nextPalindrome = string(leftHalf) + string(middleDigit) + string(reverseBytes(leftHalf))
	}

	if nextPalindrome > numStr {
		return nextPalindrome
	}

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
