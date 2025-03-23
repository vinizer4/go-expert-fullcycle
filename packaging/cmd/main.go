package main

import (
	"fmt"
	"meumodulo/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
