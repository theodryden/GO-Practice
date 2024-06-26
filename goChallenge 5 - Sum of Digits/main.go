package main

import "fmt"

func main() {

	var num int
	fmt.Println("Enter a positive integer")
	fmt.Scan(&num)

	switch {
	case num <= 0:
		fmt.Println("Please enter a positive integer")
	default:
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		fmt.Println("The sum of the digit is: %d\n", sum)
	}
}