package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

// Define MinHeap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any          { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }
func (h MinHeap) Peek() int          { return h[0] }

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	const MOD int = 1e9 + 7

	type engineer struct {
		eff int
		spd int
	}

	engineers := []engineer{}

	for i := 0; i < n; i++ {
		engineers = append(engineers, engineer{efficiency[i], speed[i]})
	}

	sort.Slice(engineers, func(i, j int) bool { return engineers[i].eff > engineers[j].eff })

	h := &MinHeap{}
	heap.Init(h)

	speedSum := 0
	maxPerf := 0

	for _, e := range engineers {
		if h.Len() == k {
			speedSum -= heap.Pop(h).(int)
		}

		heap.Push(h, e.spd)
		speedSum += e.spd

		perf := speedSum * e.eff

		if perf > maxPerf {
			maxPerf = perf
		}
	}

	return maxPerf % MOD

}

func main() {
	testCases := []struct {
		n          int
		speed      []int
		efficiency []int
		k          int
	}{
		{6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2},
		{1, []int{1}, []int{1}, 1},
		{3, []int{1, 2, 3}, []int{3, 2, 1}, 3},
		{4, []int{5, 5, 5, 5}, []int{1, 2, 3, 4}, 2},
		{5, []int{10, 1, 10, 1, 10}, []int{9, 1, 9, 1, 9}, 3},
	}

	for i, tc := range testCases {
		result := maxPerformance(tc.n, tc.speed, tc.efficiency, tc.k)
		fmt.Printf("%d.\tn: %d\n", i+1, tc.n)
		fmt.Println("\tspeed:", tc.speed)
		fmt.Println("\tefficiency:", tc.efficiency)
		fmt.Println("\tk:", tc.k)
		fmt.Println("\n\tOutput:", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
