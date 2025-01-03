package main

import (
	"fmt"
	"strings"
)

func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	points := 0

	currentSum := 0

	for i := 0; i < k; i++ {
		currentSum += calories[i]
	}

	if currentSum < lower {
		points--
	} else if currentSum > upper {
		points++
	}

	for i := k; i < len(calories); i++ {
		currentSum = currentSum - calories[i-k] + calories[i]

		if currentSum < lower {
			points--
		} else if currentSum > upper {
			points++
		}
	}

	return points

}

// Driver code
func main() {
	testCases := [][]int{
		{3, 5, 8, 2, 6},         // Test Case 1: Mixed performance
		{1, 1, 1, 1, 1},         // Test Case 2: All sums below lower limit
		{10, 12, 15, 20, 25},    // Test Case 3: All sums above upper limit
		{5, 10, 15, 20, 25, 30}, // Test Case 4: Mix of poor, normal, and good performances
		{3, 8, 7, 4, 5, 6},      // Test Case 5: Sliding window with variable performance
	}

	ks := []int{2, 2, 3, 3, 2}
	lowers := []int{7, 5, 10, 20, 7}
	uppers := []int{10, 10, 30, 40, 10}

	// Run each test case
	for i := 0; i < len(testCases); i++ {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("\tcalories = [%s]\n", joinInts(testCases[i], ", "))
		fmt.Printf("\tk = %d\n", ks[i])
		fmt.Printf("\tlower = %d\n", lowers[i])
		fmt.Printf("\tupper = %d\n", uppers[i])
		result := dietPlanPerformance(testCases[i], ks[i], lowers[i], uppers[i])
		fmt.Printf("\n\tpoints = %d\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}

func joinInts(ints []int, sep string) string {
	var strSlice []string
	for _, num := range ints {
		strSlice = append(strSlice, fmt.Sprintf(num))
	}
	return strings.Join(strSlice, sep)
}
