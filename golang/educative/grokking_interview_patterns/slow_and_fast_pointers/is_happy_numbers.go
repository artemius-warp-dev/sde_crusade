package main

import (
	"fmt"
	"strings"
)

// isHappy is the challenge function
func isHappy(num int) bool {
	slowPointer := num
	fastPointer := sumOfSquaredDigits(num)

	for fastPointer != 1 && slowPointer != fastPointer {
		slowPointer = sumOfSquaredDigits(slowPointer)
		fastPointer = sumOfSquaredDigits(sumOfSquaredDigits(fastPointer))
	}

	if fastPointer == 1 {
		return true
	}
	return false
}

// pow calculates the power of the given digit
func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

// Helper function that calculates the sum of squared digits
func sumOfSquaredDigits(number int) int {
	totalSum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}

// Driver code
func main() {
	array := []int{1, 5, 19, 25, 7}
	for i, v := range array {
		fmt.Printf("%d.\tInput number: %d\n", i+1, v)
		fmt.Printf("\n\tIs it a happy number? %t\n", isHappy(v))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
