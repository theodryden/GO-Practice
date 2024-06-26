package main

import "fmt"

func main() {

	// UFC Fight Statistics
	fighterSlice := []string{}
	opponentSlice := []string{}
	resultSlice := []string{}
	
	for {
	// Menu
	fmt.Println("Select your menu option:")
	fmt.Println("1 - Add a new fight")
	fmt.Println("2 - Remove a fight by name")
	fmt.Println("3 - list all fights for a specific fighter")
	fmt.Println("5 - Exit the program")

	// Get User Input
	var choice int
	fmt.Scan(&choice)

	// Main Program
	// Add a new fight

	var fighterName string
	var opponentName string
	var result string

	if choice == 1 {

		fmt.Print("Please enter your fighter's name:")
		fmt.Scan(&fighterName)

		fmt.Print("Please enter your opponent's name:")
		fmt.Scan(&opponentName)

		fmt.Print("Please enter the fight result:")
		fmt.Scan(&result)

		fighterSlice = append(fighterSlice, fighterName)
		opponentSlice = append(opponentSlice, opponentName)
		resultSlice = append(resultSlice, result)


	} else if choice == 2 {

		removeIndex := -1

		fmt.Print("Please enter a fighter and opponent name to remove from storage:")
		fmt.Scan(&fighterName)

		for i := range fighterSlice {
			if fighterName == fighterSlice[i] {
				removeIndex = i
			}

			}

		if removeIndex != -1 {

			fighterSlice = append(fighterSlice[:removeIndex], fighterSlice[removeIndex+1:]...)
			opponentSlice = append(opponentSlice[:removeIndex], opponentSlice[removeIndex+1:]...)
			resultSlice = append(resultSlice[:removeIndex], resultSlice[removeIndex+1:]...)
			
			fmt.Printf("Fighter: %s vs has now been removed", fighterName)
		}
			
		}
	}

	
}