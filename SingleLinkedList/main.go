package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert adds a new node at the end of the list
func (l *LinkedList) Insert(val int) {
	newNode := &Node{data: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Delete removes the node with the given value
func (l *LinkedList) Delete(val int) {
	if l.head == nil {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		return
	}
	current := l.head
	for current.next != nil && current.next.data != val {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

// Search returns the node if found, or nil
func (l *LinkedList) Search(val int) *Node {
	current := l.head
	for current != nil {
		if current.data == val {
			return current
		}
		current = current.next
	}
	return nil
}

// PrintList prints all elements in the list
func (l *LinkedList) PrintList() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)

	fmt.Println("Listed values:")
	list.PrintList()

	fmt.Println("Deleting 20...")
	list.Delete(20)
	list.PrintList()

	
	found := list.Search(35)
	if found != nil {
		fmt.Printf("Value %d found in the list.\n", found.data)
	} else {
		fmt.Printf("Value not found in the list.\n")
	}
}
