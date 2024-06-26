package main

import "fmt"

// Remove an item from a slice
func main() {
	
	// Make a slice
	slice := []string{}

	var expected string

	slice = append(slice, "Dancing")

	removeIndex := -1

	// Loop through the slice
	for i,word := range slice {
		
		if word == expected {
			removeIndex = i
		}
	}

	// Remove if not in position
	if removeIndex != -1 {

		slice = append(slice[:removeIndex], slice[removeIndex+1:]...)

		fmt.Print("The items left in slice is:", slice)

	} else {
		
	}



}