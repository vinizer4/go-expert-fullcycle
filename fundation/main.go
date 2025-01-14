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

/*
ID Example of creating a new type

	If you need to make your code more readable, you can create a new type with an underlying type.

	this example creates a new type ID with an underlying type int.
*/
type ID int // create a new type ID with underlying type int

var id ID = 42 // create a variable id of type ID

func main() {

	// var x string = "x" in local scope if the var is not used, it will cause an error
	a := "x" // this is a short variable declaration, := is used only for the first time
	// a := "y" // this will cause an error because an is already declared, use = instead
	a = "y" // this is a redeclaration

	println(a)  // "y"
	println(b)  // false
	println(c)  // 0
	println(d)  // ""
	println(e)  // 0.0
	println(h)  // "Hello, World!"
	println(id) // 42
}
