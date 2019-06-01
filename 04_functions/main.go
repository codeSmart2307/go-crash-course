package main

import "fmt"

func greeting(name string) string {
	return "Hello, " + name
}

// If all args are of the same type add it once at the end of the arguments list of the function signature
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Raneesh"))
	fmt.Println(getSum(1, 4))
}
