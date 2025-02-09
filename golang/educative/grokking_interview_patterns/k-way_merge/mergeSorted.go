package main

import (
	"fmt"
	"strings"
)

func mergeSorted(nums1 []int, m int, nums2 []int, n int) []int {
	p1 := m - 1
	p2 := n - 1

	for p := (n + m - 1); p >= 0; p-- {

		if p2 < 0 {
			break
		}

		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
	}

	return nums1
}

func main() {
	m := []int{9, 2, 3, 1, 8}
	n := []int{6, 1, 4, 2, 1}
	nums1 := [][]int{
		{23, 33, 35, 41, 44, 47, 56, 91, 105, 0, 0, 0, 0, 0, 0},
		{1, 2, 0},
		{1, 1, 1, 0, 0, 0, 0},
		{6, 0, 0},
		{12, 34, 45, 56, 67, 78, 89, 99, 0},
	}
	nums2 := [][]int{
		{32, 49, 50, 51, 61, 99},
		{7},
		{1, 2, 3, 4},
		{-99, -45},
		{100},
	}

	for index, value := range m {
		fmt.Printf("%d.\tnums1: %s, m: %d\n", index+1, strings.Replace(fmt.Sprint(nums1[index]), " ", ", ", -1), value)
		fmt.Printf("\tnums2: %s, n: %d\n", strings.Replace(fmt.Sprint(nums2[index]), " ", ", ", -1), n[index])
		result := mergeSorted(nums1[index], value, nums2[index], n[index])
		fmt.Printf("Merged list: %s\n", strings.Replace(fmt.Sprint(result), " ", ", ", -1))
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
