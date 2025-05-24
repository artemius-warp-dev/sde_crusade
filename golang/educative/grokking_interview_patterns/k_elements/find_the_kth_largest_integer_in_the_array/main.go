package main

import (
	"container/heap"
	"fmt"
	"math/big"
	"strings"
)

// MinHeap implementation for *big.Int
type MinHeap []*big.Int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Cmp(h[j]) < 0 } // Min-heap condition based on big.Int comparison
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push an element into the heap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*big.Int))
}

// Pop an element from the heap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthLargestInteger(nums []string, k int) string {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, num := range nums {
		val := new(big.Int)
		val, _ = val.SetString(num, 10)
		heap.Push(minHeap, val)

		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	return (*minHeap)[0].String()
}

// Driver code
func main() {
	testCases := []struct {
		nums []string
		k    int
	}{
		{[]string{"3", "6", "7", "10"}, 4},
		{[]string{"2", "21", "12", "1"}, 3},
		{[]string{"0", "0"}, 2},
		{[]string{"100", "200", "300", "400"}, 2},
		{[]string{"10", "100", "1000", "10000"}, 1},
	}

	for i, tc := range testCases {
		nums, k := tc.nums, tc.k
		result := kthLargestInteger(nums, k)

		fmt.Printf("%d.\t nums: %s, k: %d\n", i+1, fmt.Sprintf("[%s]", strings.Join(nums, ", ")), k)
		fmt.Printf("\t kth largest integer: %s\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
