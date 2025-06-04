package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Stack struct {
	top *Node
}

func (s*Stack)Push(val int){
	newNode:=&Node{data: val,next: s.top}
	s.top=newNode
	fmt.Println("pushed :",val)
}
func (s *Stack)Pop(){
	if s.top==nil{
		fmt.Println("Stack is empty.......!!!")
		return
	}
	fmt.Println("popped :",s.top.data)
	s.top=s.top.next
}
func (s*Stack) Peek(){
	if s.top==nil{
		fmt.Println("Stack is empty.......!!!")
		return
	}
	fmt.Println("Peek :",s.top.data)
}
func (s*Stack)Display(){
	if s.top==nil{
		fmt.Println("Stack is empty.......!!!")
		return
	}
	for temp:=s.top ; temp !=nil ;temp=temp.next{
		fmt.Println("Current stack :",temp.data)
	}
}
func main() {
	var stack Stack
	var Choice,val int

	for {
		fmt.Println("\n-----Welcome-----")
		fmt.Println("1. Push ")
		fmt.Println("2. Pop ")
		fmt.Println("3. Peek ")
		fmt.Println("4. Display ")
		fmt.Println("5. Exit ")
		fmt.Println("select your choice :")
		fmt.Scan(&Choice)
		switch Choice{
		case 1:
			fmt.Println("Enter value to push :")
			fmt.Scan(&val)
			stack.Push(val)
		case 2:
			stack.Pop()
		case 3:
			stack.Peek()
		case 4:
			stack.Display()
		case 5:
			fmt.Println("Exit succesfully...!!")
			return
		default:
			fmt.Println("Invalid charactor...!!!")
		}
	}
}