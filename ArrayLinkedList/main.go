package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (l *LinkedList) ChangeArray(array []int) {
	for _, val := range array {
		newNode := &Node{data: val}
		if l.head == nil {
			l.head = newNode
		} else {
			current := l.head
			for current.next != nil {
				current = current.next
			}
			current.next = newNode
		}
	}
}
func (l *LinkedList) Display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}
func main() {
	var list LinkedList
	Array := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Array Declared :", Array)
	list.ChangeArray(Array)
	fmt.Println("Change to Linked List...")
	list.Display()
}
