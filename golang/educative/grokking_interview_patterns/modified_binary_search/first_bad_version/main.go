package main

import (
	"fmt"
	"strings"
)

type BadVersion struct {
	bad int
}

func (bv *BadVersion) isBadVersion(version int) bool {
	return version >= bv.bad
}

type Solution struct {
	BadVersion
}

func (s *Solution) firstBadVersion(n int) int {
	first := 1
	last := n

	for first <= last {
		mid := first + (last-first)/2

		if s.isBadVersion(mid) {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return first
}

// Driver code
func main() {
	versions := []int{6, 8, 9, 11, 8}
	badVersions := []int{3, 5, 1, 11, 4}

	for i := 0; i < len(versions); i++ {
		solution := Solution{BadVersion{badVersions[i]}}
		fmt.Printf("%d.\tNumber of versions: %d\n", i+1, versions[i])
		fmt.Printf("\n\tFirst bad version: %d\n", solution.firstBadVersion(versions[i]))
		fmt.Println(strings.Repeat("-", 100))
	}
}
