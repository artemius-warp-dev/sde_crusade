package main

import (
	"fmt"
	"sort"
	"strings"
)

func kthSmallestProduct(nums1, nums2 []int, k int64) int64 {
	var lo, hi int64 = -10000000000, 10000000000
	for lo < hi {
		mid := lo + (hi-lo)/2
		if countLessOrEqual(nums1, nums2, mid) >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func countLessOrEqual(nums1, nums2 []int, val int64) int64 {
	var count int64
	for _, num := range nums1 {
		if num > 0 {
			threshold := float64(val) / float64(num)
			count += int64(bisectRight(nums2, threshold))
		} else if num < 0 {
			threshold := float64(val) / float64(num)
			count += int64(len(nums2) - bisectLeft(nums2, threshold))
		} else {
			if val >= 0 {
				count += int64(len(nums2))
			}
		}
	}
	return count
}

func bisectLeft(arr []int, target float64) int {
	return sort.Search(len(arr), func(i int) bool {
		return float64(arr[i]) >= target
	})
}

func bisectRight(arr []int, target float64) int {
	return sort.Search(len(arr), func(i int) bool {
		return float64(arr[i]) > target
	})
}

func main() {
	nums1s := [][]int{
		{1, 2, 3},
		{-3, -1, 2},
		{0, 0, 0},
		{-5, -3, -1, 2, 4},
		{-100000, 100000},
	}
	nums2s := [][]int{
		{1, 2, 3},
		{-2, 1, 3},
		{0, 0, 0},
		{1, 2, 3},
		{-100000, 100000},
	}
	ks := []int64{5, 4, 1, 8, 3}

	for i := 0; i < 5; i++ {
		result := kthSmallestProduct(nums1s[i], nums2s[i], ks[i])
		fmt.Printf("%d.\tInput array: %v\n", i+1, nums1s[i])
		fmt.Printf("\tTarget: %v, k = %d\n", nums2s[i], ks[i])
		fmt.Printf("\tResult: %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
