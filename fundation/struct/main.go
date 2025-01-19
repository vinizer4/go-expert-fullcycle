package main

import "fmt"

// Client is an example of a struct

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	wesley := Client{
		Name:   "Wesley",
		Age:    23,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", wesley.Name, wesley.Age, wesley.Active)
}
