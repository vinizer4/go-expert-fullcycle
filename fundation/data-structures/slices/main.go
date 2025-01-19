package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 50, 60, 70, 80, 90, 100} // slice literal syntax is like an array literal without the length
	// a slice is a dynamically-sized, flexible view into the elements of an array

	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)

	fmt.Printf("len=%d cap=%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	slice = append(slice, 110) // append returns a new slice with the element added

	// when append is called the capacity of the slice is doubled if the new length exceeds the capacity
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}
