package main

import "fmt"

func greeting(name string) string {
	return "Hello, " + name
}

// If all args are of the same type add it once at the end of the arguments list of the function signature
func getSum(num1, num2 int) int {
	return num1 + num2
}

// Go has support for returning multiple values from functions
func getSumAndDifference(num1, num2 int) (int, int) {
	// Multiple values can be returned from the same return statement by separating the values by commas
	return num1 + num2, num1 - num2
}

func main() {
	fmt.Println(greeting("Raneesh"))
	fmt.Println(getSum(1, 4))

	// The return values from getSumAndDifference have to be captured by 2 separate variables
	sum, difference := getSumAndDifference(5, 2)
	fmt.Println(sum, difference)
}
