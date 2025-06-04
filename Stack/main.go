package main

import "fmt"

type Stack struct{
	items[]int
}

func (S*Stack)Push(Val int){
	S.items=append(S.items, Val)
	fmt.Println("Pushed :",Val)
}
func (S *Stack)Pop(){
	if len(S.items)==0{
		fmt.Println("Stack is empty...!!")
		return
	}
	top:=S.items[len(S.items)-1]
	S.items=S.items[:len(S.items)-1]
	fmt.Println("Pop value :",top)
}
func (S*Stack)Peek(){
	if len(S.items)==0{
		fmt.Println("Stack is empty...!!")
		return
	}
	top:=S.items[len(S.items)-1]
	fmt.Println("Top value is ",top)
}
func (S *Stack)Display(){
	fmt.Println("Current stack :",S.items)
}
func main() {
var Choice ,val int
var stack Stack
	for {
		fmt.Println("-----Welcome-------")
		fmt.Println("1. Push to stack")
		fmt.Println("2. Pop to stack")
		fmt.Println("3. Peek to stack")
		fmt.Println("4. Display to stack")
		fmt.Println("5. Exit from stack")
		fmt.Println("Enter your choice :")
		fmt.Scan(&Choice)
		switch Choice{
		case 1:
			fmt.Println("Enter data to push")
			fmt.Scan(&val)
			stack.Push(val)
		case 2:
			stack.Pop()
		case 3:
			stack.Peek()
		case 4:
			stack.Display()
		case 5:
			fmt.Println("Exit sucessfully....!!!!")
			return
		default:
			fmt.Println("Invalid charactor")
		}
	}
}