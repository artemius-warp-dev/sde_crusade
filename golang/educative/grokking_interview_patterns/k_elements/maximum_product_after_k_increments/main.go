package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumProduct(nums []int, k int) int {
	const MOD int = 1e9 + 7
	minHeap := &IntHeap{}
	heap.Init(minHeap)

	for _, num := range nums {
		heap.Push(minHeap, num)

	}

	for i := 0; i < k; i++ {
		smallest := heap.Pop(minHeap).(int)
		heap.Push(minHeap, smallest+1)

	}

	result := 1
	for minHeap.Len() > 0 {
		result = (result * heap.Pop(minHeap).(int))

	}
	return result

}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 3, 3, 3, 0}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{0, 0, 0, 0, 0}, 10},
		{[]int{1, 5, 1, 1}, 3},
		{[]int{2, 2, 2, 2}, 4},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tnums = %v, k = %d\n", i+1, tc.nums, tc.k)
		result := maximumProduct(tc.nums, tc.k)
		fmt.Printf("\tMaximum Product: %d\n", result)
		fmt.Println("-------------------------------------------------------------------")

	}
}
