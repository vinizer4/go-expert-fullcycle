package main

import "fmt"

func main() {
	var myArray [3]int // array of 3 integers
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	fmt.Println(len(myArray) - 1) // 2 (3 - 1) the last index of the array

	for index, value := range myArray {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
