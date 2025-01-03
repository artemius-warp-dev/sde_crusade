package main

import (
	"fmt"
	"strings"
)

func totalFruit(fruits []int) int {
	baskets := make(map[int]int)
	collected := 0
	left := 0

	for right := 0; right < len(fruits); right++ {
		baskets[fruits[right]]++

		for len(baskets) > 2 {
			baskets[fruits[left]]--

			if baskets[fruits[left]] == 0 {
				delete(baskets, fruits[left])
			}

			left++
		}

		if collected < right-left+1 {
			collected = right - left + 1

		}
	}
	return collected

}

// Driver code
func main() {
	fruits := [][]int{
		{3, 4, 2, 1, 3, 2},
		{2, 2, 2, 3, 1, 2, 4, 4, 4, 4},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 3, 4, 1, 3, 3, 1, 2, 3, 4, 1, 5, 2, 5, 7, 7},
		{5, 4, 3, 2, 1, 1},
	}

	for i, fruitArr := range fruits {
		fmt.Printf("%d.\tFruits: [%s]\n", i+1, strings.Trim(strings.Replace(fmt.Sprint(fruitArr), " ", ", ", -1), "[]"))
		fmt.Printf("\n\tMaximum number of fruit(s) collected: %d\n", totalFruit(fruitArr))
		fmt.Println(strings.Repeat("-", 100))
	}
}
