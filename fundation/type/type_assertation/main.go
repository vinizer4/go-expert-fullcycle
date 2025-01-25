package main

import "fmt"

func main() {
	var myVar interface{} = "Wesley Willians"
	println(myVar.(string))

	result, ok := myVar.(int)
	fmt.Printf("The value of res is %v and the result of the ok is %v", result, ok)

	// this will throw a panic because the type of myVar is string
	res2 := myVar.(int)
	fmt.Printf("The value of res2 is %v", res2)
}
