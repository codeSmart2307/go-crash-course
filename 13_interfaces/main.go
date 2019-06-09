package main

import (
	"fmt"
	"math"
)

// Interfaces are NOT generic types - Other languages like Java & C# have generic types but Go, famously, does not
// Interfaces are IMPLICIT - We do not manually say that our custom type implements/satisfies some interface
// Interfaces are a contract to help us manage types - If our custom type's implementation of a function is broken
// then interfaces won't help us

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
