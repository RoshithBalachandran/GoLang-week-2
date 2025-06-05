package main

import "fmt"

type DNode struct{
	data int
	prev *DNode
	next *DNode
}
type DoubleLinkedList struct{
	head *DNode
}

func (d *DoubleLinkedList) Insert(val int){
	newNode:=&DNode{data: val}

	if d.head==nil{
		d.head =newNode
		return
	}
	current:=d.head
	for current.next !=nil{
		current=current.next
	}
	current.next=newNode
	newNode.prev=current
}

func (d *DoubleLinkedList)Delete(val int){
	current:=d.head
	
	if current==nil{
		fmt.Println("list is empty...!!!")
		return
	}
	if current.data==val{
		d.head=current.next
		if d.head!=nil{
			d.head.prev=nil
		}
		return
	}
	for current !=nil && current.data !=val{
		current=current.next
	}
	if current ==nil{
		fmt.Println("value not found..!!")
		return
	}
	if current.prev !=nil{
		current.prev.next=current.next
	}
	if current.next !=nil{
		current.next.prev=current.prev
	}
}

func (d *DoubleLinkedList) Forward(){
	current:=d.head
	fmt.Println("Forward Printing...")
	for current !=nil{
		fmt.Print(current.data,"<->")
		current=current.next
	}
	fmt.Println("nil")
}
func (d *DoubleLinkedList) Reverse(){
	current:=d.head

	if current==nil{
		fmt.Println("Task is empty...")
		return
	}
	for current.next!=nil{
		current=current.next
	}
	fmt.Println("Reversed :")
	for current !=nil{
		fmt.Print(current.data,"<->")
		current=current.prev
	}
	fmt.Println("nil")
}

func main() {
	var val int
	var n int
	var d DoubleLinkedList
	var De int
	fmt.Println("Welcome to Double Linked List")

	fmt.Println("How many number of data you want to add :")
	fmt.Scan(&n)
	fmt.Println("Enter ",n," number of data")
	for i:=0;i<n;i++{
		fmt.Scan(&val)
		d.Insert(val)
	}
	d.Reverse()
	d.Forward()

	fmt.Println("enter value to delete..!!")
	fmt.Scan(&De)
	d.Delete(De)

	d.Reverse()
	d.Forward()

}