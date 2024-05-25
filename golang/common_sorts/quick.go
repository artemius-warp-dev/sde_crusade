package main

import "fmt"

func quickSort(arr []int) {

	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)/2]

	left := make([]int, 0)
	right := make([]int, 0)
	equal := make([]int, 0)

	for _, num := range arr {

		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			equal = append(equal, num)
		}

	}

	quickSort(left)
	quickSort(right)

	copy(arr, left)
	copy(arr[len(left):], equal)
	copy(arr[len(left)+len(equal):], right)

}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array: ", arr)
	quickSort(arr)
	fmt.Println("Sorted array: ", arr)
}
