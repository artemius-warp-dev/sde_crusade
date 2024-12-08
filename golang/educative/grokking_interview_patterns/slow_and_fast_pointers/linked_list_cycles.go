package main

//import "fmt"

func detectCycle(head *EduLinkedListNode) bool {

	slowPtr, fastPtr := head, head

	for {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next
		if fastPtr == nil || fastPtr.next == nil {
			return false
		} else {
			fastPtr = fastPtr.next
		}

		//fmt.Printf("%s, %s\n", slowPtr, fastPtr)
		if slowPtr == fastPtr {
			return true
		}
	}

	return false
}
