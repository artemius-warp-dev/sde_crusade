package main

import (
	"fmt"
	"strings"
)

func removeElement(nums []int, val int) int {
	k := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[k] = nums[j]
			k++
		}
	}
	return k
}

// Driver code
func main() {
	numsArr := [][]int{
		{5, 8, 8, 5, 3},
		{50, 49, 48, 47, 46, 45},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{10, 20, 30, 40, 50},
		{0, 50},
	}

	valArr := []int{5, 48, 0, 25, 0}

	for i := 0; i < len(numsArr); i++ {
		fmt.Printf("%d.\tnums: %v\n", i+1, numsArr[i])
		fmt.Printf("\tval: %d\n", valArr[i])
		k := removeElement(numsArr[i], valArr[i])
		fmt.Printf("\tk: %d\n", k)
		fmt.Println(strings.Repeat("-", 100))
	}
}
