package main

import "fmt"

func main() {
	
	var itemPrices = []int{30,20,10,4,5}

	// calculate total cost of items in itemPrices
	var totalCost int
	

	// loop through the array
	for i := 0; i < len(itemPrices); i++ {
		totalCost += itemPrices[i]
	}
	
	// print the result

	fmt.Println(totalCost)
}

/* 
Lessons:
1. You can just create new files instead of folders for Homework


*/ 