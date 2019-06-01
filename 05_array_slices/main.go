package main

import "fmt"

func main() {
	// Arrays

	// Declare array
	var fruitArr [2] string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])

	// Declare and instantiate
	fruitArr2 := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr2)
	fmt.Println(fruitArr2[0])

	// Dynamic arrays
	fruitSlice := []string{"Apple", "Orange", "Kiwi", "Grapes"}

	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[2])
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1: 3])

}
