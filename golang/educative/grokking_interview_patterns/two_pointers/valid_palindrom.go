package main

func isPalindrome(inputString string) bool {
	start := 0
	end := len(inputString) - 1
	mid := end / 2
	for {
		if start == mid || end == mid {
			return true
		}
		if inputString[start] == inputString[end] {
			start++
			end--
		} else {
			return false
		}
	}
}
