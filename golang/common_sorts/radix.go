package main

import (
	"fmt"
	"math"
)

func radixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	max := findMax(arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		countSort(arr, exp)
	}

}

func countSort(arr []int, exp int) {

	n := len(arr)

	count := make([]int, 10)
	output := make([]int, n)

	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 10
		count[digit]++
	}

	// calculate cummulative sum
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}

	copy(arr, output)

}

func findMax(arr []int) int {
	max := math.MinInt32

	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	return max

}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array: ", arr)
	radixSort(arr)
	fmt.Println("Sorted array: ", arr)
}
