package main

import "fmt"

func findMax(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {

	nums := []int{10, 80, 60, 70, 90, 100}
	Max := findMax(nums)
	fmt.Println("Max :",Max)
}