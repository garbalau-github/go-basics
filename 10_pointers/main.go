package main

import "fmt"

func main() {
	// Points to a memory address of a value
	a := 5
	b := &a
	
	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // *int

	// Read value from memory address
	fmt.Println(*b) // 5

	// Change val with pointer
	*b = 10
	fmt.Println(a) // 10
}