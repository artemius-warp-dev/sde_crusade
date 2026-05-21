package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

// Min-heap will store candidate intervals that could cover current query:
// (interval_size, interval_right_endpoint)
type Item struct {
	size int
	r    int
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].size != h[j].size {
		return h[i].size < h[j].size
	}
	return h[i].r < h[j].r
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)   { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	type Q struct {
		q   int
		idx int
	}

	qs := make([]Q, 0, len(queries))
	for i, q := range queries {
		qs = append(qs, Q{q: q, idx: i})
	}
	sort.Slice(qs, func(i, j int) bool {
		return qs[i].q < qs[j].q
	})

	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}

	h := &MinHeap{}
	heap.Init(h)

	i := 0
	n := len(intervals)

	for _, pair := range qs {
		q := pair.q
		idx := pair.idx

		for i < n && intervals[i][0] <= q {
			left := intervals[i][0]
			right := intervals[i][1]
			size := right - left + 1
			heap.Push(h, Item{size: size, r: right})
			i++
		}

		for h.Len() > 0 && (*h)[0].r < q {
			heap.Pop(h)
		}

		if h.Len() > 0 {
			ans[idx] = (*h)[0].size
		}

	}

	return ans

}

func intervalsToString(intervals [][]int) string {
	parts := make([]string, 0, len(intervals))
	for _, in := range intervals {
		parts = append(parts, fmt.Sprintf("[%d, %d]", in[0], in[1]))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func main() {
	intervalsList := [][][]int{
		{{1, 4}, {2, 4}, {3, 6}, {4, 4}},
		{{2, 3}, {2, 5}, {1, 8}, {20, 25}},
		{{5, 10}},
		{{1, 10}, {2, 3}, {4, 8}, {6, 6}},
		{{1, 2}, {4, 4}, {7, 9}},
	}

	queriesList := [][]int{
		{2, 3, 4, 5},
		{2, 19, 5, 22},
		{4, 5, 7, 10, 11},
		{1, 2, 6, 9, 10},
		{1, 3, 4, 8, 10},
	}

	for t := 0; t < len(intervalsList); t++ {
		// Create a copy so the original test data stays unchanged after sorting
		intervals := make([][]int, len(intervalsList[t]))
		for k := range intervalsList[t] {
			intervals[k] = []int{intervalsList[t][k][0], intervalsList[t][k][1]}
		}

		queries := append([]int(nil), queriesList[t]...)

		result := minInterval(intervals, queries)

		fmt.Printf("%d.\tintervals: %s\n", t+1, intervalsToString(intervalsList[t]))
		fmt.Printf("\tqueries: %v\n", queriesList[t])
		fmt.Printf("\toutput: %v\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
