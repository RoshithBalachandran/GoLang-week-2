package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p)
	*p=45
	fmt.Println(*p)
}