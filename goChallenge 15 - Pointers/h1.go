package main

import "fmt"

/*

Description:
Write a function in Go that increments the value of an integer 
by a given amount using pointers.

Requirements:

Define a function Increment(ptr *int, amount int).
Use a pointer to modify the value of the integer.
Test your function with different initial values and increments.

*/

// Increment function
func Increment(ptr* int, amount int){

	// Store the value from the pointer
	*ptr += amount

}

func main() {

	ptr := 45
	amount := 10

	// call the pointer value
	Increment(&ptr,amount)

	// Print it out
	fmt.Println("Pointer:",ptr, "| Increment Amount:", amount)

}