package main

import (
	"fmt"
	t "time"
)

func main() {

	// var exit bool = true
	var selector int
	var exit bool = true

	// Menu Options For User
	menu := "Menu Options: Enter the number corresponding with your choice!\n1: Print Hello, World!\n2: Print the current date and time\n3: Exit the program"

	for exit {
	fmt.Println(menu)

	// Get user input
	fmt.Scan(&selector)

	// Handling the conditions
	switch selector {
	case 1:
		fmt.Println("Hello World!")
	case 2:
		fmt.Println("The time us currently:", t.Now())
	case 3:
		fmt.Println("You have chosen to exit the program.")
		exit = false
	}
}
}