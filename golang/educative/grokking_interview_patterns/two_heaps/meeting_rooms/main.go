package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

// MinHeap structure initialization
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of MinHeap given their indices
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap function swaps the value of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
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

// Pair struct to hold the end time and room number
type Pair struct {
	endTime int64
	room    int
}

// PairHeap is a MinHeap of Pairs
type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}
func (h PairHeap) Empty() bool {
	return len(h) == 0
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].endTime < h[j].endTime || (h[i].endTime == h[j].endTime && h[i].room < h[j].room)
}
func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mostBooked(meetings [][]int, rooms int) int {

	count := make([]int, rooms)

	available := &MinHeap{}
	heap.Init(available)

	usedRooms := &PairHeap{}
	heap.Init(usedRooms)

	for i := 0; i < rooms; i++ {
		heap.Push(available, i)
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for i := 0; i < len(meetings); i++ {
		startTime := int64(meetings[i][0])
		endTime := int64(meetings[i][1])

		for !usedRooms.Empty() && (*usedRooms)[0].endTime <= startTime {
			room := heap.Pop(usedRooms).(Pair).room
			heap.Push(available, room)
		}

		if available.Len() == 0 {
			pair := heap.Pop(usedRooms).(Pair)
			endTime = pair.endTime + (endTime - startTime)
			heap.Push(available, pair.room)
		}

		room := heap.Pop(available).(int)
		heap.Push(usedRooms, Pair{endTime: endTime, room: room})
		count[room]++

	}
	maxMeetingsRoom := 0
	for i := 0; i < rooms; i++ {
		if count[i] > count[maxMeetingsRoom] {
			maxMeetingsRoom = i
		}
	}

	return maxMeetingsRoom

}

// Driver code
func main() {
	meetings := [][][]int{
		{{0, 10}, {1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}},
		{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
		{{1, 2}, {0, 10}, {2, 3}, {3, 4}},
		{{0, 2}, {1, 2}, {3, 4}, {2, 4}},
		{{1, 9}, {2, 8}, {3, 7}, {4, 6}, {5, 11}},
	}
	rooms := []int{3, 3, 2, 4, 3}

	for i, meeting := range meetings {
		fmt.Printf("%d. \tMeetings: %v\n", i+1, meeting)
		fmt.Printf("\tRooms: %d\n", rooms[i])
		bookedRooms := mostBooked(meeting, rooms[i])
		fmt.Printf("\n\tRoom that held the most meetings: %d\n", bookedRooms)
		fmt.Println(strings.Repeat("-", 100))
	}
}
