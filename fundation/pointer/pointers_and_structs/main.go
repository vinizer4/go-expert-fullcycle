package main

import "fmt"

//type Client struct {
//	name string
//}
//
//func (c Client) walked() {
//	c.name = "Wesley Willians"
//	fmt.Printf("The client %v walked", c.name)
//}
//
//func main() {
//	wesley := Client{name: "Wesley"}
//	wesley.walked()
//	fmt.Printf("The value of struct with name %v", wesley.name)
//}

type Account struct {
	balance float64
}

// NewAccount this is a constructor of the Account struct with default values and return a pointer to the memory address of the struct
func NewAccount() *Account {
	return &Account{balance: 0}
}

func (a Account) simulate(value float64) float64 {
	a.balance += value
	return a.balance
}

func (a *Account) deposit(value float64) {
	fmt.Println("Depositing: ", value)
	a.balance += value
}

func main() {
	account := Account{balance: 100}
	simulatedBalance := account.simulate(200)
	fmt.Println("Simulated balance: ", simulatedBalance)
	fmt.Println("Real balance: ", account.balance)
	account.deposit(400)
	fmt.Println("Real balance after deposit: ", account.balance)

	// if you need change the real values of the struct you need to use a pointer

	accountPointer := NewAccount()
	fmt.Println("Built account with pointer: ", *accountPointer)
}
