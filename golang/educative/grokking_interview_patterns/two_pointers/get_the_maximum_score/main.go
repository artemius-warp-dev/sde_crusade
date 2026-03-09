package main

import (
	"fmt"
	"strings"
)

func maxSum(nums1 []int, nums2 []int) int {
	pointer1, pointer2 := 0, 0
	sumPath1, sumPath2 := 0, 0
	MOD := 1_000_000_007
	len1, len2 := len(nums1), len(nums2)

	for pointer1 < len1 || pointer2 < len2 {
		if pointer1 < len1 && (pointer2 == len2 || nums1[pointer1] < nums2[pointer2]) {
			sumPath1 += nums1[pointer1]
			pointer1++
		} else if pointer2 < len2 && (pointer1 == len1 || nums1[pointer1] > nums2[pointer2]) {
			sumPath2 += nums2[pointer2]
			pointer2++
		} else {
			common := nums1[pointer1]
			maxSum := sumPath1
			if sumPath2 > maxSum {
				maxSum = sumPath2
			}

			sumPath1 = maxSum + common
			sumPath2 = sumPath1
			pointer1++
			pointer2++
		}
	}

	if sumPath1 > sumPath2 {
		return sumPath1 % MOD
	}
	return sumPath2 % MOD
}

// Driver code
func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{2, 4, 5, 8, 10}, []int{4, 6, 8, 9}},      // common: 4, 8 → max = 30
		{[]int{1, 3, 5, 7, 9}, []int{3, 5, 100}},        // common: 3, 5 → max = 109
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}},   // no common → max = 40
		{[]int{2, 5, 7, 11, 13}, []int{1, 4, 6, 8, 10}}, // no common → max = 38
		{[]int{1, 2, 4, 6, 8}, []int{2, 4, 6, 7, 9}},    // common: 2, 4, 6 → max = 32
	}

	for i, tc := range testCases {
		result := maxSum(tc.nums1, tc.nums2)
		fmt.Printf("%d.\tnums1: %v\n\tnums2: %v\n\tresult: %d\n", i+1, tc.nums1, tc.nums2, result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
