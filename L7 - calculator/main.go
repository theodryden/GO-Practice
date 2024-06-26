package main

import "fmt"

func calculateAverage() float32{
	var numbers = []float32{}
	var number float32
	var running bool = true
	var sum float32

	for running == true {

		fmt.Println("Please enter a number:")
		fmt.Scan(&number)

		numbers = append(numbers, number)

		if number < 0 {
			running = false
		} else {
			numbers = append(numbers, number)
		}

		// calculate the sum

		for _,number := range numbers{
			sum += number
		}
	}
	// Calculate the average
	return sum / float32(len(numbers))
	
}

func performCalculation(n1 float32, n2 float32, o string) float32{
	var result float32
	if o == "+" {
		fmt.Println("The result is", n1 + n2)
	} else if o == "-" {
		fmt.Println("The result is", n1 - n2)
	} else if o == "x" {
		fmt.Println("The result is", n1 * n2)
	} else if o == "/" {
		if n2 != 0 {
			fmt.Println("The result is", n1 / n2)
		} else {
			fmt.Println("It is not possible to divide a number by zero")
		}
	}
	return result
}

func main() {

	var number1 float32
	var number2 float32
	var operator string
	var result float32
	var choice int
	var running bool = true

	for running {
		// Menu
		fmt.Println("1 - Perform calculation")
		fmt.Println("2 - average operation")
		fmt.Println("3 - exit")
		
		fmt.Println("Please enter an option")
		fmt.Scan(&choice)

		if choice == 1 {

			fmt.Println("Please enter the first number:")
			fmt.Scan(&number1)

			fmt.Println("Please enter the second number:")
			fmt.Scan(&number2)

			fmt.Println("Please enter the operator:")
			fmt.Scan(&operator)
			
			result = performCalculation(number1,number2,operator)

			fmt.Println("The result is", result)

		} else if choice == 2 {

			result := calculateAverage()
			fmt.Println("The average is", result)

		} else if choice == 3 {

		}

	}
	
}