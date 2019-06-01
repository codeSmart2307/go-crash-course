package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key-values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Jeff"] = "jeff@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and add key-values
	emails2 := map[string]string{"Bob":"bob@gmail.com", "Jeff":"jeff@gmail.com"}
	fmt.Println(emails2)
}

