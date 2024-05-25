package main

// Двусвязный список

// Len() // длинна списка
// First() // первый Item
// Last() // последний Item
// PushFront(v interface{}) // добавить значение в начало
// PushBack(v interface{}) // добавить значение в конец
// Remove(i Item) // удалить элемент


import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
	prev  *Node
}

type DList struct {
	head *Node
}

func (list *DList) Len() int {
	len := 0
	current := list.head
	for current != nil {
		len++
		current = current.next
	}
	return len
}

func (list *DList) isItem(item interface{}) bool {
	cur := list.head
	for cur != nil {
		if cur.value == item {
			return true
		}
		cur = cur.next
	}
	return false
}

func (list *DList) First() Node {
	return *list.head
}

func (list *DList) Last() Node {
	last := list.head
	for last.next != nil {
		last = last.next
	}
	return *last
}

func (list *DList) Remove(item interface{}) {

	// Check if the list is empty
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	// If the item to be removed is at the head of the list
	if list.head.value == item {
		list.head = list.head.next
		if list.head != nil {
			list.head.prev = nil
		}
		return
	}

	current := list.head

	for current != nil {
		if current.value == item {
			current.prev.next = current.next
			current.next.prev = current.prev
			return
		}
		current = current.next
	}
	fmt.Println("NOT FOUND")
}

func (list *DList) PushFront(v interface{}) {
	newNode := &Node{value: v, next: nil, prev: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.next = list.head
	list.head.prev = newNode
	list.head = newNode
}

func (list *DList) PushBack(v interface{}) {

	newNode := &Node{value: v, next: nil, prev: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	last := list.head
	for last.next != nil {
		last = last.next
	}

	last.next = newNode
	newNode.prev = last

}

func (list *DList) traverse() {
	current := list.head
	for current != nil {
		fmt.Printf("%v ", current.value)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := DList{head: nil}

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushFront(4)

	list.traverse()
	fmt.Println(list.Len())
	fmt.Println(list.First())
	fmt.Println(list.Last())
	list.Remove(2)
	list.Remove(4)
	list.traverse()

}
