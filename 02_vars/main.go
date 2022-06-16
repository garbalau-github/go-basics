package main

import "fmt"

func main() {
	var age int32 = 24
	const isCool = true

	// Shorthand
	name := "Nick"
	size := 1.3

	email, pass := "nick@gmail.com", "qwerty"

	// Printing
	fmt.Println(name, age, isCool, email, pass)
	
	// Get Type
	fmt.Printf("%T\n", size)
}