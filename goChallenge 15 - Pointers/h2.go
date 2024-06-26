package main

import "fmt"

/*

Description:
Write a function in Go that 
doubles the value of an integer using a pointer.

Requirements:

Define a function Double(ptr *int).
Use a pointer to double the value of the integer.
Test your function with different initial values.

*/


func Double(ptr* int){
	
	*ptr *= 2

}

func main() {

	// Test 1 ptr := 10
	// Test 2 ptr := 100
	// Test 3 ptr := 1000

	ptr := 1000

	fmt.Println("Initial Value:", ptr)

	Double(&ptr)

	fmt.Println("Doubled Value:", ptr)

}