package main

import (
	"fmt"
	"math"
)

// Interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// Method for Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Method for Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// General
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	// Interface as data types representing a set of 
	// method signatures for structs

	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}