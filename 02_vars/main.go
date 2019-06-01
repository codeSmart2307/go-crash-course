package main

import "fmt"

func main() {
	// Using var
	var name = "Raneesh"
	fmt.Println(name)
	fmt.Printf("%T\n", name) // Print formatting verbs - https://golang.org/pkg/fmt/

	// Using :=
	// This type of variable instantiation cannot be used in the global scope, only in the function scope
	// You cannot instantiate static typed variables using this method
	name2 := "Suleka"
	size := 1.3
	age, email := 21, "raneesh2307@gmail.com" // Multiple variables can be instantiated in a single statement
	fmt.Println(name2, age, email, size)
	fmt.Printf("%T\n", size)
}
