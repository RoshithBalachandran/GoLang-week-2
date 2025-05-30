package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println("value of p:",p)
	*p=20
	fmt.Println("Value changed using pointer",*p)


	nums:=make([]int,5)
	nums[0]=10
	nums[1]=20
	nums[2]=30
	fmt.Println(nums)
}