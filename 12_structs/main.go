package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	firstName, lastName, city, gender string
	age																int
}

// Methods: value receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " and I'm a " + strconv.Itoa(p.age)
}

// Methods: pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return 
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init Person using struct
	person1 := Person{
		firstName: "Samantha",
		lastName: "Smith",
		city: "Boston",
		gender: "F",
		age: 25,
	}

	// Alternative
	person2 := Person{
		"Bob",
		"Johnson",
		"Boston",
		"M",
		30,
	}

	fmt.Println(person1)
	fmt.Println(person2)

	// Get single thing
	fmt.Println(person1.firstName)
	
	// Raise up age a bit
	for i := 0; i < 10; i++ {
		person1.age++
	}
	fmt.Println(person1.age);

	// Call methods
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Williams")
	person2.getMarried("Thomson")
	fmt.Println(person2.lastName)
	fmt.Println(person1.lastName)
}