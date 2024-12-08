package main

import "fmt"

func nextPoint(nums []int, i int, steps int) int {
	step := 0
	if steps > 0 {
		for step != steps {

			i++
			step++
			if i > len(nums)-1 {
				i = 0
			}
		}
	} else {
		steps = steps * -1
		for step != steps {
			i--
			step++
			if i < 0 {
				i = len(nums) - 1
			}
		}

	}
	return i
}

func isChangedDirection(nums []int, slow int, fast int) bool {
	if nums[slow] > 0 && nums[fast] > 0 || nums[slow] < 0 && nums[fast] < 0 {
		return false
	} else {
		return true
	}

}

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {

		slow, fast := i, i
		fmt.Println(i)

		for {
			slow = nextPoint(nums, slow, nums[slow])
			fast = nextPoint(nums, fast, nums[fast])
			if isChangedDirection(nums, slow, fast) {
				break
			}
			newFast := nextPoint(nums, fast, nums[fast])
			if newFast == fast {
				break
			} else {
				fast = newFast
			}

			if isChangedDirection(nums, slow, fast) {
				break
			}

			fmt.Println(i, slow, fast)
			if slow == fast {
				return true
			}
		}

	}

	return false
}

func main() {
	//test := []int{1, 4, 3, 2, 1}
	//test := []int{1, 3, -2, -4, 1}
	//test := []int{2, 1, -1, -2}
	//test := []int{5, 4, -2, -1, 3}
	//test := []int{-4737, -4455}
	test := []int{1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1, 1, -1, 4, -3, -1}
	/*
	   slow 3 fast 1(4) 1(0)
	   slow 1 (4)  3(1)
	   slow 1(0)
	   slow 1 fast 3  1


	*/
	fmt.Println(circularArrayLoop(test))
}
