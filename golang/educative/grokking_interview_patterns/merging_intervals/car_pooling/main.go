package main

import (
	"fmt"
	"strings"
)

func carPooling(trips [][]int, capacity int) bool {
	timestamp := make([]int, 1001)
	for _, trip := range trips {
		numPassengers := trip[0]
		start := trip[1]
		end := trip[2]

		timestamp[start] += numPassengers
		timestamp[end] -= numPassengers
	}

	usedCapacity := 0
	for _, passengerChange := range timestamp {
		usedCapacity += passengerChange
		if usedCapacity > capacity {
			return false
		}
	}

	return true
}

// Helper function to print a trip array nicely
func formatTrips(trips [][]int) string {
	parts := make([]string, len(trips))
	for i, trip := range trips {
		parts[i] = fmt.Sprintf("[%d, %d, %d]", trip[0], trip[1], trip[2])
	}
	return strings.Join(parts, ", ")
}

// Driver code
func main() {
	testCases := []struct {
		trips    [][]int
		capacity int
	}{
		{[][]int{{2, 1, 5}, {3, 3, 7}}, 4},
		{[][]int{{2, 1, 5}, {3, 3, 7}}, 5},
		{[][]int{{3, 2, 6}, {1, 4, 7}, {2, 5, 8}}, 5},
		{[][]int{{1, 0, 4}, {2, 2, 6}, {3, 5, 8}}, 6},
		{[][]int{{4, 1, 5}, {1, 3, 7}, {2, 6, 8}}, 5},
	}

	for i, tc := range testCases {
		fmt.Printf("%d.\tInput: trips = [%s], capacity = %d\n", i+1, formatTrips(tc.trips), tc.capacity)
		result := carPooling(tc.trips, tc.capacity)
		fmt.Printf("\tCan complete all trips? %v\n", result)
		fmt.Println(strings.Repeat("-", 100))
	}
}
