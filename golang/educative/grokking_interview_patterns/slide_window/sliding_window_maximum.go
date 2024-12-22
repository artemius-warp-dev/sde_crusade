package main

import "fmt"

func getMaxFromWindow(nums []int, start, end int) (int, bool) {
  if end > len(nums) {
    return -1, false
  }
  window := nums[start:end]
  max := window[0]
  for _, n := range window {
    if n > max {
      max = n
    }
  }
  return max, true
} 

func findMaxSlidingWindow(nums []int, w int) []int {
	result := make([]int, 0)
  
  for i, _ := range nums {
    max, bound := getMaxFromWindow(nums, i, w + i)
    fmt.Println(max)
    if bound {
      result = append(result, max) 
    }
  }
	return result
}
