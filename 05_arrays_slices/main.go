package main

import "fmt"

func main() {

	// Arrays (have to be typed and fixed length)
	fruitArr := [2]string{"Apple", "Orange"}

	// Slice (array without fixed length)
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitArr)
	fmt.Println(fruitSlice)

	// Get length of array / slice
	fmt.Println(len(fruitArr))

	// Get element in range
	fmt.Println(fruitSlice[1:2])
}