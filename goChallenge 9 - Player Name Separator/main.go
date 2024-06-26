package main

import (
	"fmt"
	s "strings"
)

func main() {

	var firstName,lastName string
	pName := "Michael Jordan"

	// turn the name into an array so the index can be accessed in a for loop
	splitName := s.Fields(pName)

	// write a loop to go through the name and if it's a certain index, split into first and last name
	for i, name := range splitName {
		// if iterator number is 0
		if i == 0 {
			firstName = name
		} else {
			lastName = name
			
		}
	}
	fmt.Printf("First Name: %v\n",firstName)
	fmt.Printf("Last Name: %v\n",lastName)

}