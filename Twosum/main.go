package main

import "fmt"

func Check(ar []int, tar int) []int {

	n := len(ar)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if ar[i]+ar[j] == tar {
				return []int{ar[i], ar[j]}
			}
		}
	}
	return nil
}

func main() {
	array := []int{1, 3, 4, 5, 8, 9}
	target := 10
	result := Check(array, target)
	if result != nil {
		fmt.Printf("The Pair of two number sum is %d + %d = %d", result[0], result[1], target)
	} else {
		fmt.Println("No Pair found in the array")
	}
}
