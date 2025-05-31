package main

import "fmt"

func main(){

	array:=[5]int{1,2,3,4,5}
	n:=len(array)
	for i := n; i >= 0; i-- {
		fmt.Println(array[i])
	}
	
}