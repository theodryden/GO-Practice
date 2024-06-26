package main

import "fmt"

/*

Description:
Write a function in Go that 
appends an integer to a slice using a pointer to the slice.

Requirements:

Define a function AppendToSlice(slice *[]int, value int).
Use a pointer to the slice to append the new value.
Test your function by appending values to slices of different lengths.

*/

func AppendToSlice(slice *[]int, value int){

	 *slice = append(*slice, value)

}

func main() {
	// 1. Test Slice = []int{5}
	// 2. Test Slice = []int{2,4,6,8,10,12}
	// 3. Test Slice = []int{10,100,1000,10000,100000}

	slice := []int{10,100,1000,10000,100000}

	fmt.Println("Old Slice:", slice)

	// 1. Test Numbers = 5
	// 2. Test Numbers = 10
	// 3. Test Numbers = 1000000

	AppendToSlice(&slice, 1000000)

	fmt.Println("New Slice:", slice)
}