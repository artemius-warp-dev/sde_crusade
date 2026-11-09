package main

import (
	"fmt"
	"strings"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2

	for left <= right {
		mid := left + (right-left)/2
		midSquared := mid * mid

		if midSquared == x {
			return mid
		} else if midSquared < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right

}

func main() {
	testCases := []int{0, 1, 25, 15, 2147395600}

	for idx, x := range testCases {
		result := mySqrt(x)
		fmt.Printf("%d.\tInput: x = %d\n", idx+1, x)
		fmt.Printf("\tResult: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
