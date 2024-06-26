package main

import "fmt"

func main() {
		/* 
	Task: Calculate the average temperature 
	from an array of daily temperatures 
	and print the result.
	*/

	dailyTemp := [5]int{3, 1, 2, 4, 5}

	// Variables 
	sum := 0
	
	/* loop through the numbers
	for _, temp := range dailyTemp {
		sum += temp
	} */

	// looping another way
	for i := 0; i < len(dailyTemp); i++ {
		sum += dailyTemp[i]
	}

	// calculate the average of sum
	average := sum / len(dailyTemp)

	// Print
	fmt.Println("The sum of the numbers in the array is:",sum)
	fmt.Println("The average of the numbers in the array is:",average)

}




