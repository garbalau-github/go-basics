package main

import "fmt"

func main() {
	// Map - key & value pairs

	// Define map
	emails := make(map[string]string) // [key]value 

	// Assign key & values
	emails["bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"]) // Get value by key

	// Delete from map
	delete(emails, "mike")

	fmt.Println(emails)

}