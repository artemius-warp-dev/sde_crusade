package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// MinHeap structure initialization
type MinHeap []int

// newMinHeap function initializes an instance of MinHeap
func newMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

// Len function returns the length of MinHeap
func (h MinHeap) Len() int {
	return len(h)
}

// Empty function returns true if empty, false otherwise
func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of the MinHeap given their indexes
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap function swaps the values of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() int {
	if h.Empty() {
		return 0
	} else {
		return h[0]
	}
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	kNumbersMinHeap := newMinHeap()

	for i := 0; i < k; i++ {
		heap.Push(kNumbersMinHeap, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > kNumbersMinHeap.Top() {
			heap.Pop(kNumbersMinHeap)
			heap.Push(kNumbersMinHeap, nums[i])
		}
	}

	return kNumbersMinHeap.Top()

}

// Driver code
func main() {
	inputList := [][]int{
		{1, 5, 12, 2, 11, 9, 7, 30, 20},
		{5, 2, 9, -3, 7},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 4, 6, 0, 2},
		{3, 5, 2, 3, 8, 5, 3},
	}
	k := []int{3, 1, 9, 1, 4}

	for i, list := range inputList {
		fmt.Printf("%d.\tInput array: %s \n\tValue of k: %d\n", i+1, strings.Replace(fmt.Sprint(list), " ", ", ", -1), k[i])
		fmt.Printf("\tkth largest number: %d\n", findKthLargest(list, k[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
