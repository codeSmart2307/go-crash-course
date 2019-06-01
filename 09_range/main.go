package main

import "fmt"

func main() {
	ids := []int{33, 45, 67, 78, 22, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID %d\n", i, id)
	}

	// Not using index
	// Use an underscore instead of a variable if you're not planning to use the index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add IDs together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum of IDs: ", sum)

	// Range with map
	emails := map[string]string{"Bob":"bob@gmail.com", "Jeff":"jeff@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Extract just the key
	for k := range emails {
		fmt.Printf("Name: " + k + "\n")
	}
}
