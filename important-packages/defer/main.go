package main

import "fmt"

func main() {
	fmt.Println("First line")
	// defer keyword is used to delay the execution of a statement until the surrounding function returns.
	defer fmt.Println("Second line")
	fmt.Println("Third line")
}
