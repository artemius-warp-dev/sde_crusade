package main

import (
	"container/heap"
	"strconv"
)

type IntervalHeap [][]int

func (h IntervalHeap) Len() int {
	return len(h)
}

func (h IntervalHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h IntervalHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntervalHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *IntervalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntervalHeap) Peek() []int {
	return (*h)[0]
}

type Interval struct {
	Start int
	End   int
}

func (i *Interval) IntervalInit(start int, end int) {
	i.Start = start
	i.End = end

}

func (i *Interval) str() string {
	out := "(" + strconv.Itoa(i.Start) + ", " + strconv.Itoa(i.End) + ")"
	return out
}




func employeeFreeTime(schedule [][]*Interval) []*Interval {
	h := &IntervalHeap{}
	heap.Init(h)

	for i:=0; i< len(schedule); i++ {
		employeeSchedule := schedule[i]
		interval := employeeSchedule[0]
		heap.Push(h, []int{interval.Start, i, 0})
	}

	result := []*Interval{}

	previous := (*h).Peek()[0]

	for h.Len() > 0 {
		tuple := heap.Pop(h).([]int)
		i := tuple[1]
		j := tuple[2]

		interval := schedule[i][j]
		if interval.Start > previous {
			result = append(result, &Interval{Start: previous, End: interval.Start})
		}

		previous = max(previous, interval.End)

		if j+1 < len
	}
}
