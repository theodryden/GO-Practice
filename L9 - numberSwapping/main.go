package main

import "fmt"

/*

*/
// You need pointers to reference the function outside the main() 
// Every time you need to change a variable, slice, array, outside the main

// Catching variables in a function, you need pointers
// Must put the pointer behind the parameter * 

// So it's like * means store the memory and & means access the memory

func swap(pointer1* int, pointer2* int){

	// Access the pointer parameters in front * next to the variables
	z := *pointer1
	*pointer1 = *pointer2
	*pointer2 = z 

}

func main() {
	x := 5
	y := 7

	swap(&x, &y)

	fmt.Println("x:", x, "-", "y", y)

}