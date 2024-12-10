package main

import "fmt"

func traverse_until_meet(nums []int, slowPtr, fastPtr, slowStep, fastStep int) int {

	for {

		//if slowPtr == len(nums)-1 {
		//	return -1
		//}
		fmt.Println(nums[slowPtr], nums[fastPtr])
		if nums[slowPtr] == nums[fastPtr] {
			return slowPtr
		}

		slowPtr = slowPtr + slowStep
		fastPtr = fastPtr + fastStep

		if fastPtr > len(nums)-1 {
			fastPtr = (0 + fastStep) - 1
		}

		if slowPtr > len(nums)-1 {
			slowPtr = (0 + slowStep) - 1
		}
	}
}

func findDuplicate(nums []int) int {

	fast, slow := 2, 0

	firstPoint := traverse_until_meet(nums, slow, fast, 1, 2)
	if firstPoint == -1 {
		return -1
	}

	fmt.Println(firstPoint, "first point")

	//secondPoint := traverse_until_meet(nums, 0, firstPoint, 1, 1)
	secondPoint := -1
	if secondPoint == -1 {
		return -1
	}

	fmt.Println(secondPoint, "second point")

	return secondPoint

}

func main() {

	data := []int{1, 5, 4, 3, 2, 4, 6}
	/*

	   1 4
	   5 2
	   4 6
	   3 5
	   2 3
	   4 4


	   1 4
	   5 6
	   4 1
	   3 5
	   2 4
	   4 3
	   6 2
	   1 4
	   5 6
	   4 1
	   3 5
	*/

	fmt.Println(findDuplicate(data))

}
