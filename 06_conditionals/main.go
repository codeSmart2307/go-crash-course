package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	}

	a := 10
	b := 10

	// If-else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", a, b)
	} else {
		fmt.Printf("%d is less than %d\n", a, b)
	}

	// Else-if
	color := "blue"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is invalid")
	}

	// Switch
	switch color {
		case "red":
			fmt.Println("color is red")
		case "blue":
			fmt.Println("color is blue")
		default:
			fmt.Println("color is invalid")
	}
}
