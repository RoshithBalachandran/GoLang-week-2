package main

import "fmt"

type Queue struct{
	items [] int
}

func (q *Queue)Enqueue(val int){
	q.items=append(q.items, val)
	fmt.Println("Value Enqueue :",val)
}
func (q *Queue)Denqueue(){
	if len(q.items)==0{
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println("Enqueue :",q.items[0])
	q.items=q.items[1:]
	
}

func main() {
	var Choice,val int
	var q Queue
	for {
		fmt.Println("Welcome to Queue")
		fmt.Println("1. Enqueue")
		fmt.Println("2. Dequeue")
		fmt.Println("3. Display")
		fmt.Println("4. Exit")
		fmt.Scan(&Choice)
		switch Choice{
		case 1:
			fmt.Println("Enter value to Enqueue")
			fmt.Scan(&val)
			q.Enqueue(val)
		case 2:
			q.Denqueue()
		case 3:
			fmt.Println("Current queue :",q.items)
		case 4:
			fmt.Println("Exit sucesssfully...!!")
			return
		default:
			fmt.Println("Invalid charactor..!!!")
		}
	}
}