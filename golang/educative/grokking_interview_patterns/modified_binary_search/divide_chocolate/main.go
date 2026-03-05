package main

import (
	"fmt"
	"strings"
)

func canDivide(sweetness []int, k, minSweetness int) bool {
	totalSweetness := 0
	pieces := 0
	for _, sweet := range sweetness {
		totalSweetness += sweet
		if totalSweetness >= minSweetness {
			pieces++
			totalSweetness = 0
		}
	}
	return pieces >= k+1
}

func maximizeSweetness(sweetness []int, k int) int {

	total := 0
	for _, sweet := range sweetness {
		total += sweet
	}

	low, high := 1, total/(k+1)
	result := low

	for low <= high {
		mid := (low + high) / 2
		if canDivide(sweetness, k, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return result

}

// Driver code
func main() {
	testCases := []struct {
		sweetness []int
		k         int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
		{[]int{5}, 0},
		{[]int{1, 2, 2, 1, 2, 2, 1}, 3},
		{[]int{1, 1, 1, 1, 1, 1, 1}, 6},
		{[]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, 20},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tsweetness = %v\n\tk = %d\n", i+1, tc.sweetness, tc.k)
		result := maximizeSweetness(tc.sweetness, tc.k)
		fmt.Printf("\n\toutput: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
