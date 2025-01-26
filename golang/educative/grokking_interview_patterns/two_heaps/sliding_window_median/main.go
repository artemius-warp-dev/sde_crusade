package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Structure for MaxHeap
type MaxHeap []int

// newMaxHeap function intializes an instance of the MaxHeap
func newMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

// Len function returns the length of the MaxHeap
func (h MaxHeap) Len() int {
	return len(h)
}

// Empty returns true if the MaxHeap is empty, false otherwise
func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

// Greater returns true if the first of the given elements is greater than the second one
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap function swaps the values at the given indices
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() int {
	if h.Empty() {
		return 0
	} else {
		return h[0]
	}
}

// Push function inserts the element in the MaxHeap
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// printMedians is a helper function to print the medians list
func printMedians(medians []float64) {
	fmt.Printf("[")
	for i, v := range medians {
		fmt.Printf("%.1f", v)
		if i != len(medians)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n\n")
}

func medianSlidingWindow(nums []int, k int) []float64 {
	medians := []float64{}
	outgoingNum := make(map[int]int)
	smallList := new(MaxHeap)
	largeList := new(MinHeap)

	for i := 0; i < k; i++ {
		heap.Push(smallList, nums[i])
	}

	for i := 0; i < k/2; i++ {
		heap.Push(largeList, heap.Pop(smallList))

	}
	balance := 0
	i := k

	for {
		if k%2 == 0 {
			medians = append(medians, (float64(smallList.Top())+float64(largeList.Top()))/2)
		} else {
			medians = append(medians, (float64(smallList.Top())))
		}

		if i >= len(nums) {
			break
		}

		outNum := nums[i-k]
		inNum := nums[i]

		i++

		if outNum <= smallList.Top() {
			balance--
		} else {
			balance++
		}

		if _, ok := outgoingNum[outNum]; ok {
			outgoingNum[outNum]++
		} else {
			outgoingNum[outNum] = 1
		}

		if !smallList.Empty() && inNum <= smallList.Top() {
			balance++
			heap.Push(smallList, inNum)
		} else {
			balance--
			heap.Push(largeList, inNum)
		}

		if balance < 0 {
			heap.Push(smallList, largeList.Pop())
		} else if balance > 0 {
			heap.Push(largeList, smallList.Pop())
		}

		balance = 0

		for smallList.Len() > 0 && outgoingNum[smallList.Top()] > 0 {
			outgoingNum[heap.Pop(smallList).(int)]--

		}

		for largeList.Len() > 0 && outgoingNum[largeList.Top()] > 0 {
			outgoingNum[heap.Pop(largeList).(int)]--

		}

	}
	return medians

}

// Driver code
func main() {
	inputList := [][]int{
		{3, 1, 2, -1, 0, 5, 8},
		{1, 2},
		{4, 7, 2, 21},
		{22, 23, 24, 56, 76, 43, 121, 1, 2, 0, 0, 2, 3, 5},
		{1, 1, 1, 1, 1},
	}
	k := []int{4, 1, 2, 5, 2}

	for index, value := range inputList {
		fmt.Printf("%d.\tInput array: %s, k = %d\n", index+1, strings.Replace(fmt.Sprint(value), " ", ", ", -1), k[index])
		fmt.Printf("\tMedians: ")
		printMedians(medianSlidingWindow(value, k[index]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
