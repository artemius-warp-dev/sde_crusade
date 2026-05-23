package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// A custom type to represent a max-heap of class data based on gain
type ClassHeap struct {
	data []classData
}

type classData struct {
	gain   float64
	passes int
	total  int
}

func (h ClassHeap) Len() int           { return len(h.data) }
func (h ClassHeap) Less(i, j int) bool { return h.data[i].gain > h.data[j].gain }
func (h ClassHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *ClassHeap) Push(x interface{}) {
	h.data = append(h.data, x.(classData))
}

func (h *ClassHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func gain(passes, total int) float64 {
	return float64(passes+1)/(float64(total+1)) - float64(passes)/float64(total)
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h := &ClassHeap{}
	heap.Init(h)

	for _, cls := range classes {
		passes := cls[0]
		total := cls[1]
		g := gain(passes, total)
		heap.Push(h, classData{gain: g, passes: passes, total: total})
	}

	for i := 0; i < extraStudents; i++ {
		top := heap.Pop(h).(classData)
		top.passes++
		top.total++
		top.gain = gain(top.passes, top.total)
		heap.Push(h, top)
	}

	var totalRatio float64
	for h.Len() > 0 {
		top := heap.Pop(h).(classData)
		totalRatio += float64(top.passes) / float64(top.total)
	}

	return totalRatio / float64(len(classes))
}

func main() {
	classes := [][][]int{
		{{1, 2}, {3, 5}, {2, 2}},
		{{2, 4}, {3, 9}, {4, 5}, {2, 10}},
		{{1, 3}, {2, 4}, {3, 6}},
		{{5, 10}, {2, 3}, {3, 7}, {4, 8}},
		{{10, 20}, {5, 5}, {8, 12}, {6, 15}},
	}

	extraStudents := []int{2, 4, 3, 5, 3}

	for i := 0; i < len(classes); i++ {
		fmt.Printf("%d.\tClasses: ", i+1)
		for _, cls := range classes[i] {
			fmt.Printf("[%d, %d] ", cls[0], cls[1])
		}
		fmt.Printf("\n\tExtra Students: %d\n", extraStudents[i])

		result := maxAverageRatio(classes[i], extraStudents[i])
		fmt.Printf("\n\tFinal Average Pass Ratio: %.6f\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
