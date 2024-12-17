package main

import "fmt"

func SplitCircularLinkedList(head *EduLinkedListNode) []*EduLinkedListNode {
	//size := 0
	slow, fast := head, head.next.next
	for fast != head {
		slow = slow.next
		fast = fast.next
		if fast != head {
			fast = fast.next
		} else {
			break
		}

	}

	head1, head2 := head, slow.next
	slow.next = head1
	ptr := head2
	for ptr.next != head {
		ptr = ptr.next
	}
	ptr.next = head2
	result := []*EduLinkedListNode{head1, head2}
	fmt.Printf("%+v\n", head1)
	fmt.Printf("%+v\n", head2)

	return result
}
