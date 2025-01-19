package main

import "fmt"

func main() {

	salaries := map[string]float64{"Wesley": 1000.0, "João": 2000.0, "Maria": 3000.0}

	fmt.Println(salaries)         // map[João:2000 Maria:3000 Wesley:1000]
	fmt.Println(salaries["João"]) // 2000

	delete(salaries, "Wesley") // delete the key "Wesley" from the map
	fmt.Println(salaries)      // map[João:2000 Maria:3000]

	salaries["Wes"] = 5000.0 // add a new key to the map

	fmt.Println(salaries) // map[João:2000 Maria:3000 Wes:5000]

	sal := make(map[string]float64) // create a map with make
	sal1 := map[string]float64{}    // create a map with a composite literal

	fmt.Println(sal)  // map[]
	fmt.Println(sal1) // map[]

	// Iterating over a map
	for name, salary := range salaries {
		fmt.Printf("Name: %s, Salary: %.2f\n", name, salary)
	}

	// blank identifier
	for _, salary := range salaries {
		fmt.Printf("Salary: %.2f\n", salary)
	}
}
