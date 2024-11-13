package main

func isPalindrome(inputString string) bool {
	start := 0
	end := len(inputString) - 1
	for start < end {
		if inputString[start] != inputString[end] {
			return false
		}
		start++
		end--
	}

	return true
}
