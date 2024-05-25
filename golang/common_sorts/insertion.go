package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = key
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array: ", arr)
	insertionSort(arr)
	fmt.Println("Sorted array: ", arr)
}
