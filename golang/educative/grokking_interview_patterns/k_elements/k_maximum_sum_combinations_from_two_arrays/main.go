package main


import (
    "fmt"
	"container/heap"
	"sort"
	"strings"
)


type Pair struct {
	sum, i, j int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].sum > h[j].sum } // Max heap (largest sum at the top)
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func maxCombinations(arr1 []int, arr2 []int, k int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(arr1)))
	sort.Sort(sort.Reverse(sort.IntSlice(arr2)))

	h := &MaxHeap{}
	heap.Init(h)

	visited := make(map[[2]int]bool)

	heap.Push(h, Pair{arr1[0] + arr2[0], 0,0})
	visited[[2]int{0,0}] = true

	var result []int

	for count := 0; count < k; count++ {
		top := heap.Pop(h).(Pair)
		sum, i,j := top.sum, top.i, top.j
		result = append(result, sum)

		if i+1 < len(arr1) && !visited[[2]int{i+1,j}] {
			heap.Push(h, Pair{arr1[i+1] + arr2[j], i+1, j})
			visited[[2]int{i+1,j}] = true 
		}

		if j+1 < len(arr2) && !visited[[2]int{i, j+1}] {
			heap.Push(h, Pair{arr1[i] + arr2[j+1], i, j+1})
			visited[[2]int{i,j+1}] = true
		}
	}
	return result
}


func main() {
	testCases := []struct {
		arr1 []int
		arr2 []int
		k    int
	}{
		{[]int{1, 4, 2}, []int{3, 6, 5}, 3},
		{[]int{10, 15, 30}, []int{20, 25, 10}, 2},
		{[]int{1, 1, 1}, []int{1, 1, 1}, 2},
		{[]int{5, 7}, []int{8, 3}, 1},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 3},
	}

	for i, testCase := range testCases {
		arr1 := testCase.arr1
		arr2 := testCase.arr2
		k := testCase.k

		fmt.Printf("%d.\t arr1: %v, arr2: %v, k: %d\n", i+1, arr1, arr2, k)

		result := maxCombinations(arr1, arr2, k)

		fmt.Printf("\n\t Top %d Maximum Sums: %v\n", k, result)
		fmt.Println(strings.Repeat("-", 100))
	}
}


