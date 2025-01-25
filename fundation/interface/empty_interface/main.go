package main

import "fmt"

func main() {

	// this is the same of a generics in other languages like Java, C# or C++
	var x interface{} = 10
	var y interface{} = "Hello, World"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("The type of var is %T and the value is %v\n", t, t)
}
