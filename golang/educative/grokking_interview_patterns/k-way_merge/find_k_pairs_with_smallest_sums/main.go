package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func kSmallestPairs(list1 []int, list2 []int, k int) [][]int {
	minHeap := newMinHeap()
	pairs := make([][]int, 0)

	for i, _ := range list1 {
		heap.Push(minHeap, Set{sum: list1[i] + list2[0], i: i, j: 0})
	}

	counter := 1

	for (!minHeap.Empty()) && (counter <= k) {
		poppedValue := heap.Pop(minHeap).(Set)
		i, j := poppedValue.i, poppedValue.j

		pairs = append(pairs, []int{list1[i], list2[j]})
		nextElement := j + 1

		if len(list2) > nextElement {
			heap.Push(minHeap, Set{sum: list1[i] + list2[nextElement], i: i, j: nextElement})
		}
		counter++
	}
	return pairs
}

func main() {
	list1 := [][]int{
		{2, 8, 9},
		{1, 2, 300},
		{1, 1, 2},
		{4, 6},
		{4, 7, 9},
		{1, 1, 2},
	}

	list2 := [][]int{
		{1, 3, 6},
		{1, 11, 20, 35, 300},
		{1, 2, 3},
		{2, 3},
		{4, 7, 9},
		{1},
	}
	k := []int{9, 30, 1, 2, 5, 4}

	for i, list := range list1 {
		fmt.Printf("%d.\t Input Pairs: %s,%s \n\tK = %d\n", i+1, strings.Replace(fmt.Sprint(list), " ", ", ", -1), strings.Replace(fmt.Sprint(list2[i]), " ", ", ", -1), k[i])
		fmt.Printf("\tPairs with the smallest sum are: %s\n", strings.Replace(fmt.Sprint(kSmallestPairs(list, list2[i], k[i])), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
