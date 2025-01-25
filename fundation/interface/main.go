package main

import "fmt"

// Address is an example of a struct

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

/* This is an example of a struct composed of another struct
   this is called composition in Go
   Is similar to inheritance in POO (Programing Oriented Object)

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address
}
*/

// is the same of an object creation in POO (Programing Oriented Object)
// is this case the Address struct is a field of the Client struct

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address Address
}

// this is an example of a method in a struct in GO
// this method is associated with the Client struct
// is the same of a method in a class in POO (Programing Oriented Object)

// in GO the interface are implemented implicitly
// GO interface are a set of methods don't accept fields

type Pessoa interface {
	Disable()
}

type Firm struct {
	Name string
}

func (firm Firm) Disable() {
	fmt.Printf("The firm %s is disabled\n", firm.Name)
}

func (client Client) Disable() {
	client.Active = false
	fmt.Printf("The client %s is disabled\n", client.Name)
}

func disabled(pessoa Pessoa) {
	pessoa.Disable()
}

func main() {
	wesley := Client{
		Name:   "Wesley",
		Age:    23,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", wesley.Name, wesley.Age, wesley.Active)

	disabled(wesley)

	MyFirm := Firm{}

	disabled(MyFirm)
}
