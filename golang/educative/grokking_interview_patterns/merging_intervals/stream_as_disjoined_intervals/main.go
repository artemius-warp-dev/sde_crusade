package main

import (
	"fmt"
	"strings"
)

type SummaryRanges struct {
	intervals map[int]int
	starts    []int
}

func Constructor() *SummaryRanges {
	return &SummaryRanges{
		intervals: make(map[int]int),
		starts:    make([]int, 0),
	}
}

type pair struct{ first, second int }

func (sr *SummaryRanges) AddNum(value int) {
	newStart := value
	newEnd := value

	idx := upperBound(sr.starts, value)
	var nextInterval *pair
	if idx < len(sr.starts) {
		k := sr.starts[idx]
		nextInterval = &pair{first: k, second: sr.intervals[k]}
	} else {
		nextInterval = nil
	}

	var prevInterval *pair
	if idx > 0 {
		k := sr.starts[idx-1]
		prevInterval = &pair{first: k, second: sr.intervals[k]}

		if prevInterval.second >= value {
			return
		}

		if prevInterval.second == value-1 {
			newStart = prevInterval.first
		}
	} else {
		prevInterval = nil
	}

	if nextInterval != nil && nextInterval.first == value+1 {
		newEnd = nextInterval.second
		delete(sr.intervals, nextInterval.first)
		sr.starts = append(sr.starts[:idx], sr.starts[idx+1:]...)
	}

	if _, ok := sr.intervals[newStart]; ok {
		sr.intervals[newStart] = newEnd
	} else {
		pos := upperBound(sr.starts, newStart)
		sr.starts = append(sr.starts, 0)
		copy(sr.starts[pos+1:], sr.starts[pos:])
		sr.starts[pos] = newStart
		sr.intervals[newStart] = newEnd
	}

}

func (sr *SummaryRanges) GetIntervals() [][]int {
	result := make([][]int, 0, len(sr.starts))
	for _, s := range sr.starts {
		result = append(result, []int{s, sr.intervals[s]})
	}
	return result
}

func upperBound(a []int, x int) int {
	lo, hi := 0, len(a)
	for lo < hi {
		mid := (lo + hi) / 2
		if a[mid] <= x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// --- Driver (prints like the C++ program) ---

func intervalsToString(ivals [][]int) string {
	parts := make([]string, 0, len(ivals))
	for _, ab := range ivals {
		parts = append(parts, fmt.Sprintf("[%d, %d]", ab[0], ab[1]))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func main() {
	commandsList := [][]string{
		{"SummaryRanges", "addNum", "getIntervals"},
		{"SummaryRanges", "addNum", "addNum", "addNum", "getIntervals"},
		{"SummaryRanges", "addNum", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"},
		{"SummaryRanges", "addNum", "addNum", "addNum", "addNum", "addNum", "getIntervals"},
		{"SummaryRanges", "addNum", "addNum", "addNum", "addNum", "addNum", "getIntervals"},
	}

	argsList := [][][]int{
		{{}, {5}, {}},
		{{}, {1}, {3}, {2}, {}},
		{{}, {3}, {5}, {}, {2}, {}, {6}, {}},
		{{}, {1}, {4}, {2}, {9}, {3}, {}},
		{{}, {1}, {0}, {8}, {7}, {6}, {}},
	}

	for t := 0; t < len(commandsList); t++ {
		commands := commandsList[t]
		args := argsList[t]

		outputs := make([]string, 0, len(commands))
		var obj *SummaryRanges

		for i := 0; i < len(commands); i++ {
			cmd := commands[i]
			if cmd == "SummaryRanges" {
				obj = Constructor()
				outputs = append(outputs, "null")
			} else if cmd == "addNum" {
				val := args[i][0]
				obj.AddNum(val)
				outputs = append(outputs, "null")
			} else if cmd == "getIntervals" {
				res := obj.GetIntervals()
				outputs = append(outputs, intervalsToString(res))
			}
		}

		// Print commands
		fmt.Printf("%d\t [", t+1)
		for i := 0; i < len(commands); i++ {
			fmt.Printf("\"%s\"", commands[i])
			if i+1 < len(commands) {
				fmt.Printf(", ")
			}
		}
		fmt.Println("]")

		// Print args
		fmt.Print("\t [")
		for i := 0; i < len(args); i++ {
			fmt.Print("[")
			for j := 0; j < len(args[i]); j++ {
				fmt.Print(args[i][j])
				if j+1 < len(args[i]) {
					fmt.Print(", ")
				}
			}
			fmt.Print("]")
			if i+1 < len(args) {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")

		// Print outputs
		fmt.Print("\n\t Output: [")
		for i := 0; i < len(outputs); i++ {
			fmt.Print(outputs[i])
			if i+1 < len(outputs) {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
		fmt.Println(strings.Repeat("-", 100))
	}
}
