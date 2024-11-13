package main

import "fmt"

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

type TestCase struct {
	Input string
	Expected bool
}

func runTests() {
	testCases := []TestCase{
		{"", true},
		{"a", true},
		{"racecar", true},               
		{"hello", false},                
		{"level", true},                
		{"😊a😊", true},                
		{"абба", true},                 
		{"абвгд", false},               
		{"A man a plan a canal Panama", false}, 
		{"あいういあ", true},             
		{"あいうえお", false},
	}


	for i, testCase := range testCases {
		result := isPalindrome(testCase.Input)
		if result == testCase.Expected {
			fmt.Printf("Test %d passed! Input: '%s' -> Output: %v\n", i+1, testCase.Input, result) 
		} else {
			fmt.Printf("Test %d failed! Input: '%s' -> Expected: %v, Got: %v\n", i+1, testCase.Expected, result)
		}
	}
}


func main() {
	runTests()
}
