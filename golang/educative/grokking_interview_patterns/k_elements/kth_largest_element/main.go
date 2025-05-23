// Template for the min heap

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
	return h[0]
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

type KthLargest struct {
	topKHeap *MinHeap
	k        int
}

func (kl *KthLargest) KthLargestInit(k int, nums []int) {

	fmt.Println("\n\tInitializing the heap")
	kl.topKHeap = &MinHeap{}
	heap.Init(kl.topKHeap)
	kl.k = k

	fmt.Println("\tk =", kl.k)
	fmt.Println("\tHeap: []")

	for _, element := range nums {
		fmt.Println("Adding element", element, "to the heap.")
		kl.add(element)
	}

}

func (kl *KthLargest) add(val int) int {
	if kl.topKHeap.Len() < kl.k {
		heap.Push(kl.topKHeap, val)
		fmt.Println("\n\tSize of the heap is less than k, so we’ll push", val)

	} else if val > (*kl.topKHeap)[0] {
		fmt.Println("\n\tSize of the heap is equal to k and", val, "is greater than the top of heap", (*kl.topKHeap)[0])

		fmt.Println("\tWe will poll the top element from heap and offer the new element.")
		heap.Pop(kl.topKHeap)
		heap.Push(kl.topKHeap, val)
	} else {
		fmt.Println("\n\tSize of the heap is equal to k and", val, "is equal to or smaller than the top of heap.")
		fmt.Println("\tNo operation will be performed.")

	}

	kl.printHeap()
	return (*kl.topKHeap)[0]

}

func main() {
	nums := []int{3, 6, 9, 10}
	temp := []int{3, 6, 9, 10}
	fmt.Print("Initial stream: ")
	printArray(nums)
	fmt.Println("\nk:", 3)
	var kLargest KthLargest
	kLargest.KthLargestInit(3, nums)
	val := []int{4, 7, 10, 8, 15}
	for i := 0; i < len(val); i++ {
		fmt.Println("\tAdding a new number", val[i], "to the stream")
		temp = append(temp, val[i])
		fmt.Print("\tNumber stream: ")
		printArray(temp)
		fmt.Println("\n\tKth largest element in the stream:", kLargest.add(val[i]))
		fmt.Println(strings.Repeat("-", 100))
	}
}

func (kl *KthLargest) printHeap() {
	fmt.Print("\tHeap: [")
	for i := 0; i < kl.topKHeap.Len(); i++ {
		fmt.Print((*kl.topKHeap)[i])
		if i != kl.topKHeap.Len()-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func printArray(arr []int) {
	fmt.Print("[")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		if i != len(arr)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}
