package main

import (
	"fmt"
	"strings"
)

func pickMaxSubsequence(digits []int, subseqLength int) []int {
	if subseqLength == 0 {
		return []int{}
	}

	toRemove := len(digits) - subseqLength
	stack := make([]int, 0, len(digits))
	for _, digit := range digits {
		for toRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, digit)
	}
	if len(stack) > subseqLength {
		stack = stack[:subseqLength]
	}
	return stack
}

func isGreatSuffix(seq1 []int, i int, seq2 []int, j int) bool {
	for i < len(seq1) && j < len(seq2) && seq1[i] == seq2[j] {
		i++
		j++
	}
	if j == len(seq2) {
		return true
	}
	if i < len(seq1) && seq1[i] > seq2[j] {
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mergeSequence(seq1 []int, seq2 []int) []int {
	p1, p2 := 0, 0
	merged := make([]int, 0, len(seq1)+len(seq2))
	for p1 < len(seq1) || p2 < len(seq2) {
		if isGreatSuffix(seq1, p1, seq2, p2) {
			merged = append(merged, seq1[p1])
			p1++
		} else {
			merged = append(merged, seq2[p2])
			p2++
		}
	}
	return merged
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	m, n := len(nums1), len(nums2)
	minDigitsFromNums1 := max(n, k-n)
	maxDigitsFromNums1 := min(k, m)

	bestSequence := []int{}

	for digitsFromNums1 := minDigitsFromNums1; digitsFromNums1 <= maxDigitsFromNums1; digitsFromNums1++ {
		subsequence1 := pickMaxSubsequence(nums1, digitsFromNums1)
		subsequence2 := pickMaxSubsequence(nums2, k-digitsFromNums1)
		candidateSequence := mergeSequence(subsequence1, subsequence2)
		if isGreatSuffix(candidateSequence, 0, bestSequence, 0) {
			bestSequence = candidateSequence
		}
	}

	return bestSequence

}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
		k     int
	}{
		{[]int{5, 1, 0}, []int{9, 2, 3}, 4},
		{[]int{4, 6, 2}, []int{1, 7, 8, 3}, 5},
		{[]int{2, 2, 1}, []int{2, 9}, 3},
		{[]int{7, 5, 3}, []int{4, 6, 8}, 4},
		{[]int{1, 4, 9}, []int{9, 1, 4}, 5},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\t nums1 = %v, nums2 = %v, k = %d\n", i+1, tc.nums1, tc.nums2, tc.k)
		actual := maxNumber(tc.nums1, tc.nums2, tc.k)
		fmt.Printf("\t The maximum number possible is: %v\n", actual)
		fmt.Println(strings.Repeat("-", 100))
	}
}
