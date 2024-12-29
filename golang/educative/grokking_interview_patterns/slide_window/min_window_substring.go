package main

import (
	"fmt"
	"math"
	"strings"
)

func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}

	reqCount := make(map[rune]int)
	window := make(map[rune]int)

	for _, c := range t {
		reqCount[c]++
	}

	// for _, c := range t {
	// 	window[c] = 0
	// }

	current, required := 0, len(reqCount)
	//fmt.Println(required, current)

	res, resLen := [2]int{-1, -1}, math.MaxInt32

	left := 0

	for right := 0; right < len(s); right++ {
		c := rune(s[right])

		if _, ok := reqCount[c]; ok {
			window[c]++
		}

		if _, ok := reqCount[c]; ok && window[c] == reqCount[c] {
			current++
		}

		for current == required {
			if (right - left + 1) < resLen {
				res = [2]int{left, right}
				resLen = (right - left + 1)
			}

			if _, ok := reqCount[rune(s[left])]; ok {
				window[rune(s[left])]--
			}

			if _, ok := reqCount[rune(s[left])]; ok && window[rune(s[left])] < reqCount[rune(s[left])] {
				current--
			}

			left++
		}
	}
	left, right := res[0], res[1]

	if resLen != math.MaxInt32 {
		return s[left : right+1]
	} else {
		return ""
	}

}

func main() {
	s := []string{"PATTERN", "LIFE", "ABRACADABRA", "STRIKER", "DFFDFDFVD"}
	t := []string{"TN", "I", "ABC", "RK", "VDD"}

	s = append(s, "afffffdcccc")
	t = append(t, "afffccd")

	for i := 0; i < len(s); i++ {
		fmt.Println(i+1, ".\ts: ", s[i], "\n\tt: ", t[i], "\n\tThe minimum substring containing ", t[i], " is: ", minWindow(s[i], t[i]))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
