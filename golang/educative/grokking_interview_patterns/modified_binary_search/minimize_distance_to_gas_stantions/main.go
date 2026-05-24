package main

import (
	"fmt"
	"strings"
)

func minimizeGasDistance(stations []int, k int) float64 {
	left := 0.0
	right := float64(stations[len(stations)-1] - stations[0])

	epsilon := 1e-6

	for right-left > epsilon {
		mid := (left + right) / 2
		if isPossible(stations, k, mid) {
			right = mid
		} else {
			left = mid
		}
	}

	return left
}

func isPossible(stations []int, k int, distance float64) bool {
	requiredStattions := 0

	for i := 1; i < len(stations); i++ {
		gap := float64(stations[i] - stations[i-1])

		requiredStattions += int(gap / distance)

		if requiredStattions > k {
			return false
		}
	}

	return true
}

// Driver code
func main() {
	testCases := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{0, 10, 20, 30, 40, 50, 60, 70, 80, 90},
		{0, 3, 8, 14, 25, 27, 35, 47, 58, 72},
		{0, 8, 15, 37, 45, 52, 68, 95, 123, 150,
			178, 221, 260, 321, 389, 398, 412, 464, 531, 600},
		{5, 8, 15, 25, 40, 60, 68, 74, 104, 116,
			121, 130, 134, 159, 177, 179, 195, 206, 220, 233},
	}

	kValues := []int{9, 5, 3, 19, 5}

	for i := 0; i < len(testCases); i++ {
		stations := testCases[i]
		k := kValues[i]
		result := minimizeGasDistance(stations, k)
		fmt.Printf("%d:\tStations = %v, k = %d\n", i+1, stations, k)
		fmt.Printf("\n\tMinimum possible penalty = %.5f\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
