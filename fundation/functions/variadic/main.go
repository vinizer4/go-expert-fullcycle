package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5)) // 15
}

// example of a variadic function
// this function can receive any number of arguments
func sum(numbes ...int) int {
	total := 0
	for _, number := range numbes {
		total += number
	}
	return total
}
