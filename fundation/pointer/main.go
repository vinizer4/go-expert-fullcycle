package main

func main() {
	// memory -> address -> value

	a := 10
	println(a)

	// &a -> memory address of a
	println(&a)

	// & is used to get the memory address or set a pointer to a memory address
	var ponteiro *int = &a // ponteiro is a pointer to a memory address of a

	println(ponteiro)

	*ponteiro = 20
	println(a)

	b := &a
	println(*b) // * is used to get the value of the memory address

	// if you change the value of the memory address of b you change the value of a too
	// because b is a pointer to the memory address of a
	*b = 30
	println(a)
	println(*b)
}
