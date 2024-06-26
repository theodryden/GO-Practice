package main

import (
	"fmt"
)


func main() {
	var num int
	
	fmt.Println("Enter a number between 1 and 7:")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("This number corresponds with Monday")
	case 2:
		fmt.Println("This number corresponds with Tuesday")
	case 3:
		fmt.Println("This number corresponds with Wednesday")
	case 4:
		fmt.Println("This number corresponds with Thursday")
	case 5:
		fmt.Println("This number corresponds with Friday")
	case 6:
		fmt.Println("This number corresponds with Saturday")
	case 7:
		fmt.Println("This number corresponds with Sunday")
	default:
		fmt.Println("You need to choose a number between 1 and 7")
	}

}
