package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	//firstName string
	//lastName string
	//city string
	//gender string
	//age int

	// This is a shorter way to define a struct
	firstName, lastName, city, gender string
	age int
}

// Two types of methods can be defined in Go
// 1. Value receivers
// 2. Pointer receivers (can change/ mutate properties of an object)

// Greeting method (Value receiver)
// (p Person) is a receiver. The receiver is used similar to the "this" keyword in OOP languages
// In a receiver function, if we do not plan on using the receiver, we can remove the variable and just leave the
// struct type - eg: func (Person) greet() string { <logic goes here> }
func (p Person) greet() string {
	// The string converter package's Itoa() function is used to convert the age to a string here
	// You can also use the string() method which is built-in by default and needs no package imports
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// hasBirthday method (Pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (Pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return;
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Raneesh", lastName: "Gomez", city: "Ragama", gender: "M", age: 22}

	// Alternative way to initialize a person
	person2 := Person{"Ashley", "Halpe","Deniyaya", "F", 24}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person2)
	fmt.Println(person2.age)
	person2.age++
	fmt.Println(person2.age)

	fmt.Println("**********************************************************")

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())

	person2.getMarried(person1.lastName)
	fmt.Println(person2.greet())
}
