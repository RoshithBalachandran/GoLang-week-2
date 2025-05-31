package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4}
	fmt.Println("the initizialized array :", array)
	array = append(array, 5, 6, 7)
	fmt.Println("Appending values to the array :", array)
	array = append(array[:2], array[4:]...)
	fmt.Println("Deleted value from array", array)

	value := 7
	found := false
	for _, val := range array {
		if value == val {
			found = true
			break
		}
		
	}
	fmt.Println("value Found :?", found)
}
