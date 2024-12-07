package main

import (
	"strconv"
)

func squareSum(num int) int {
	digits := fetchDigits(num)
	sum := 0
	for _, digit := range digits {
		sum = sum + (digit * digit)
	}
	return sum
}

func fetchDigits(num int) []int {
	numberStr := strconv.Itoa(num)
	var digits []int
	for _, char := range numberStr {
		digit, _ := strconv.Atoi(string(char))
		digits = append(digits, digit)
	}
	return digits
}

func isHappy(num int) bool {

	slowPtr, fastPtr := num, squareSum(num)
	if num == 1 {
		return true
	}

	for fastPtr != slowPtr {
		if fastPtr == 1 {
			return true
		}
		slowPtr = squareSum(slowPtr)
		fastPtr = squareSum(squareSum(fastPtr))
	}

	return false
}
