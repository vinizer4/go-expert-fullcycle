package main

const h = "Hello, World!"

var (
	b bool    // default value is false
	c int     // default value is 0
	d string  // default value is ""
	e float64 // default value is 0.0
)

/*
Example of explicit initialization

var (
	b bool   = true
	c int    = 1
	d string = "Hello, World!"
	e float64 = 1.0
)
*/

func main() {
	// var x string = "x" in local scope if the var is not used, it will cause an error
	a := "x" // this is a short variable declaration, := is used only for the first time
	// a := "y" // this will cause an error because a is already declared, use = instead
	a = "y" // this is a redeclaration
	println(a, b)
}
