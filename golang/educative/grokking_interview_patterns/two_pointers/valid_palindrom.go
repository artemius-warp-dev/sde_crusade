package main

func isPalindrome(inputString string) bool {
	runes := []rune(inputString)
	start := 0
	end := len(runes) - 1
	for start < end {
		if runes[start] != runes[end] {
			return false
		}
		start++
		end--
	}

	return true
}
