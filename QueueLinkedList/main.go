package main

import "fmt"

type Node struct{
    data int
    next *Node
}
type Queue struct{
    head *Node
    tail *Node
}

func (q *Queue) Enqueue(val int){
    newNode:=&Node{data: val}
    if q.head==nil{
        q.head=newNode
        q.tail=newNode
    }else{
        q.tail.next=newNode
        q.tail=newNode
    }
    fmt.Println("Enqueue Added :",val)
}

func (q *Queue) Dequeue(){
    if q.head==nil{
        fmt.Println("Queue is empty...!!!")
        return
    }
    fmt.Println("Dequeue : ",q.head.data)
    q.head=q.head.next
    if q.head==nil{
        q.tail=nil
    }
}

func (q *Queue)Display(){
    if q.head==nil{
        fmt.Println("Queue is empty ...!!")
        return
    }
    current :=q.head
    fmt.Print("Current Queue :")
    for current !=nil{
        fmt.Print(current.data," ")
        current =current.next
    }
    fmt.Println("nil")
}
func main() {
    var Choice ,val int
    var q Queue
	for {
		fmt.Println("\n------------Welcome LinkedLisy queue---------")
		fmt.Println("1. Enqueuee to ququeue")
		fmt.Println("2. dequeuee to ququeue")
		fmt.Println("3. Display to ququeue")
		fmt.Println("4. Exit to ququeue")
        fmt.Println("Enter your Choice :")
        fmt.Scan(&Choice)
        switch Choice{
        case 1:
            fmt.Println("Enter value to Enqueue ")
            fmt.Scan(&val)    
            q.Enqueue(val)
        case 2:
            q.Dequeue()
        case 3:
            q.Display()
        case 4:
            fmt.Println("Exit Sucessfully...!!")
            return
        default:
            fmt.Println("Invalid Charactor..!!1")
        }
	}
}
