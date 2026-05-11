package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	revs1 := strings.Split(version1, ".")
	revs2 := strings.Split(version2, ".")

	p1 := 0
	p2 := 0

	maxLen := len(revs1)
	if len(revs2) > maxLen {
		maxLen = len(revs2)
	}

	for p1 < maxLen || p2 < maxLen {
		val1 := 0
		if p1 < len(revs1) {
			val1, _ = strconv.Atoi(revs1[p1])
		}
		val2 := 0
		if p2 < len(revs2) {
			val2, _ = strconv.Atoi(revs2[p2])
		}

		if val1 < val2 {
			return -1
		} else if val1 > val2 {
			return 1
		}

		p1 += 1
		p2 += 1
	}

	return 0

}

func main() {
	testCases := [][2]string{
		{"0.1", "1.1"},
		{"1.0.1", "1"},
		{"7.5.2.4", "7.5.3"},
		{"1.0.0", "1"},
		{"2.0.0.1", "2.0.0.2"},
	}

	for i, tc := range testCases {
		v1 := tc[0]
		v2 := tc[1]
		result := compareVersion(v1, v2)
		fmt.Printf("%d.\tInput array: [\"%s\", \"%s\"]\n", i+1, v1, v2)
		fmt.Printf("\tResult: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
