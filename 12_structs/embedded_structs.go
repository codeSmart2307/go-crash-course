package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// When we just declare but do not instantiate a variable like below, Go assigns each of them a zero value
	/*
	* string = ""
	* int, float - 0
	* bool - false
	 */
	var dilum person

	fmt.Println(dilum)
	fmt.Printf("%+v\n", dilum) // This prints out the actual values currently assigned to the struct properties

	fmt.Println("================================================================================================")

	sam := person{firstName: "Sam", lastName: "Smith", contact: contactInfo{email: "samsmith@gmail.com", zipCode: 10405}}

	sam.print()

	// This is the long way to call a pointer receiver method
	// The short way is used in the below call to updateName()
	pointerToSuleka := &sam
	pointerToSuleka.updateNameLongWay("Raneesh")
	sam.print()

	// Here we just pass the "reference to the value" to the updateName() method because Go is smart enough to know
	// to convert the value to a pointer before feeding it into the function logic
	// This way, our code will be much shorter
	sam.updateName("Dilum")
	sam.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// This is the long way to define a pointer receiver method
// The short way is defined in the updateName() method below
func (pointerToPerson *person) updateNameLongWay(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
