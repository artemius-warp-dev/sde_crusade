package main

import (
	"container/heap"
	"fmt"
	"strings"
)

type Pair struct {
	val int
	idx int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].val != h[j].val {
		return h[i].val < h[j].val
	}
	return h[i].idx < h[j].idx
}

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func NewMinHeapFrom(nums []int) *MinHeap {
	mh := &MinHeap{}
	for i, v := range nums {
		*mh = append(*mh, Pair{val: v, idx: i})
	}
	heap.Init(mh)
	return mh
}

func getFinalState(nums []int, k int, multiplier int) []int {
	if k == 0 || multiplier == 1 {
		return nums
	}

	h := NewMinHeapFrom(nums)

	for step := 0; step < k; step++ {
		elem := heap.Pop(h).(Pair)
		val, idx := elem.val, elem.idx
		newVal := val * multiplier
		nums[idx] = newVal
		heap.Push(h, Pair{val: newVal, idx: idx})
	}

	return nums

}

func arrayToString(arr []int) string {
	strs := make([]string, len(arr))
	for i, v := range arr {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return "[" + strings.Join(strs, ", ") + "]"
}

func main() {
	testCases := []struct {
		nums       []int
		k          int
		multiplier int
	}{
		{[]int{3, 1, 2}, 2, 3},
		{[]int{2, 2, 1, 1}, 3, 2},
		{[]int{4, 1, 7}, 5, 1},
		{[]int{1, 2, 3}, 5, 2},
		{[]int{4, 4, 4}, 2, 5},
	}

	for idx, tc := range testCases {
		orig := make([]int, len(tc.nums))
		copy(orig, tc.nums)

		result := getFinalState(tc.nums, tc.k, tc.multiplier)

		fmt.Printf("%d.        nums = %s\n", idx+1, arrayToString(orig))
		fmt.Printf("          k = %d\n", tc.k)
		fmt.Printf("          multiplier = %d\n\n", tc.multiplier)
		fmt.Printf("          output = %s\n", arrayToString(result))
		fmt.Println(strings.Repeat("-", 100))
	}
}
