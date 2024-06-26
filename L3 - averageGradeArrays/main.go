package main

import "fmt"


func main() {

	// arrays practice
	var numbers [5]int

	// for loop to repeatedly ask for numbers to add into the array
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Insert a num:")
		fmt.Scan(&numbers[i])
	}

	var mean float32 = 0 

	for i := 0; i < 5; i++ {
		mean = mean + float32(numbers[i])
	}

	mean /= 5 
	fmt.Println("The average is", mean)

	
	for i := 0; i < 5; i++ {
		fmt.Println("Number", i+1, ":", numbers[i])
	}
	

}


	/*
	var a [2]string
	var b [10]int

	a[0] = "hello"
	a[1] = "why"

	b[0] = 10
	b[8] = 9
	b[3] = 3
	
	fmt.Println(a)
	fmt.Println(b)
	*/