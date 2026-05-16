package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
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

	for i := 0; i < len(schedule); i++ {
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

		if j+1 < len(schedule[i]) {
			nextInterval := schedule[i][j+1]
			heap.Push(h, []int{nextInterval.Start, i, j + 1})
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function for displaying interval list
func display(l1 []*Interval) string {
	if len(l1) == 0 {
		return "[]"
	}

	resultStr := "["
	for i := 0; i < len(l1)-1; i++ {
		resultStr += "[" + strconv.Itoa(l1[i].Start) + ", "
		resultStr += strconv.Itoa(l1[i].End) + "], "
	}

	resultStr += "[" + strconv.Itoa(l1[len(l1)-1].Start) + ", "
	resultStr += strconv.Itoa(l1[len(l1)-1].End) + "]"
	resultStr += "]"

	return resultStr
}

// Driver code
func main() {
	inputs := [][][]*Interval{
		{{&Interval{1, 2}, &Interval{5, 6}}, {&Interval{1, 3}}, {&Interval{4, 10}}},
		{{&Interval{1, 3}, &Interval{6, 7}}, {&Interval{2, 4}}, {&Interval{2, 5}, &Interval{9, 12}}},
		{{&Interval{2, 3}, &Interval{7, 9}}, {&Interval{1, 4}, &Interval{6, 7}}},
		{{&Interval{3, 5}, &Interval{8, 10}}, {&Interval{4, 6}, &Interval{9, 12}}, {&Interval{5, 6}, &Interval{8, 10}}},
		{{&Interval{1, 3}, &Interval{6, 9}, &Interval{10, 11}}, {&Interval{3, 4}, &Interval{7, 12}}, {&Interval{1, 3}, &Interval{7, 10}}, {&Interval{1, 4}}, {&Interval{7, 10}, &Interval{11, 12}}},
		{{&Interval{1, 2}, &Interval{3, 4}, &Interval{5, 6}, &Interval{7, 8}}, {&Interval{2, 3}, &Interval{4, 5}, &Interval{6, 8}}},
		{{&Interval{1, 2}, &Interval{3, 4}, &Interval{5, 6}, &Interval{7, 8}, &Interval{9, 10}, &Interval{11, 12}}, {&Interval{1, 2}, &Interval{3, 4}, &Interval{5, 6}, &Interval{7, 8}, &Interval{9, 10}, &Interval{11, 12}}, {&Interval{1, 2}, &Interval{3, 4}, &Interval{5, 6}, &Interval{7, 8}, &Interval{9, 10}, &Interval{11, 12}}, {&Interval{1, 2}, &Interval{3, 4}, &Interval{5, 6}, &Interval{7, 8}, &Interval{9, 10}, &Interval{11, 12}}},
	}

	for i, schedule := range inputs {
		fmt.Printf("%d.\tEmployee Schedules:\n", i+1)
		for _, s := range schedule {
			fmt.Printf("\t\t%s\n", display(s))
		}
		fmt.Printf("\n\tEmployees' free time: %s\n", display(employeeFreeTime(schedule)))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
