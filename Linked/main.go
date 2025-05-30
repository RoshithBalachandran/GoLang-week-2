package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
	}
	current := l.head
	if current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (l *LinkedList) Display() {
	current := l.head
	if current != nil {
		fmt.Printf("%d ->",current.data)
		current=current.next
	}
	fmt.Println("nil")
}
func main() {
list:=LinkedList{}
list.Insert(10)
list.Insert(20)
list.Insert(30)
list.Display()
}