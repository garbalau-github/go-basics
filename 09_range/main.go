package main

import "fmt"

func main() {
	// Range to loop through arrays, slices, maps

	ids := []int{36, 75, 34, 51, 89, 38, 26}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Without index
	for _, id := range ids {
		fmt.Printf("id: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// Range with map
	emails := make(map[string]string)
	emails["bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["mike"] = "mike@gmail.com"

	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}
	
}