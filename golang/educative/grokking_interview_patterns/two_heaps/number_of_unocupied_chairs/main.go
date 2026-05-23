// minheap.go
package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

type Pair struct {
	leavingTime int
	chairNumber int
}

type MinHeap []Pair

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].leavingTime < h[j].leavingTime }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	heap.Init(h)
	return h
}

func smallestChair(times [][]int, targetFriend int) int {
	sortedFriends := make([][3]int, len(times))
	for i, time := range times {
		sortedFriends[i] = [3]int{i, time[0], time[1]}
	}

	sort.Slice(sortedFriends, func(i, j int) bool {
		return sortedFriends[i][1] < sortedFriends[j][1]
	})

	availableChairs := []int{}
	occupiedChairs := &MinHeap{}
	heap.Init(occupiedChairs)

	chairIndex := 0

	for _, friend := range sortedFriends {
		friendID, arrival, leaving := friend[0], friend[1], friend[2]

		for occupiedChairs.Len() > 0 && (*occupiedChairs)[0].leavingTime <= arrival {
			feedChair := heap.Pop(occupiedChairs).(Pair).chairNumber
			availableChairs = append(availableChairs, feedChair)
			sort.Ints(availableChairs)
		}

		var assignedChair int
		if len(availableChairs) > 0 {
			assignedChair = availableChairs[0]
			availableChairs = availableChairs[1:]
		} else {
			assignedChair = chairIndex
			chairIndex++
		}

		heap.Push(occupiedChairs, Pair{leaving, assignedChair})

		if friendID == targetFriend {
			return assignedChair
		}

	}
	return -1
}

// Driver code
func main() {
	testCases := [][][]int{
		{{3, 6}, {1, 6}, {4, 5}, {2, 4}, {5, 7}},
		{{3, 5}, {2, 6}, {1, 7}},
		{{5, 10}, {2, 3}, {3, 8}, {1, 6}},
		{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
		{{1, 10}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
	}
	targetFriends := []int{4, 0, 3, 2, 4}

	for i, times := range testCases {
		result := smallestChair(times, targetFriends[i])
		fmt.Printf("%d.\t Times: %v \n\t Target friend: %d \n\n\t Chair number: %d\n", i+1, times, targetFriends[i], result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
