package main

import "fmt"

func main() {
	pointer:=10
	fmt.Println("Before Modifing using pointer",pointer)
	after:=&pointer
	*after=25
	fmt.Println("After changine value using Pointer",*after)
}