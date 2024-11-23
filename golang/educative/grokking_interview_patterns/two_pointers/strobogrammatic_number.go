package main

func isStrobogrammatic(num string) bool {
	// Define the valid strobogrammatic mappings
	strobogrammaticMap := map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}

	// Initialize two pointers
	left, right := 0, len(num)-1

	// Move the pointers towards the center
	for left <= right {
		// Get the digits from both ends
		leftChar := num[left]
		rightChar := num[right]

		// Check if the current digits are valid and match each other's rotated version
		if val, ok := strobogrammaticMap[leftChar]; !ok || val != rightChar {
			return false
		}

		// Move the pointers
		left++
		right--
	}

	// If all pairs are valid
	return true
}
