package main

func main() {
	a := 1
	b := 2

	// in GO not exist else if statement, instead use else and if in the same line
	if a < b {
		println("a is less than b")
	} else {
		println("a is not less than b")
	}

	// operators comparison in GO
	// == equal
	// != not equal
	// < less than
	// <= less than or equal
	// > greater than
	// >= greater than or equal
	// && and
	// || or
	// ! not

	// if with a short statement
	if c := 3; c < 4 {
		println("c is less than 4")
	}

	// switch statement
	switch a {
	case 1:
		println("a is 1")
	case 2:
		println("a is 2")
	default:
	}
}
