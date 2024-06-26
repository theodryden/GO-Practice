package main

import "fmt"

// Add new player name, shirt number and goal numbers

func showMenu(){
	fmt.Println("1 - Add new player")
	fmt.Println("2 - Remove player")
	fmt.Println("3 - List all players")
	fmt.Println("4 - Show player with most goals")
	fmt.Println("5 - Exit")

}

func addPlayer(playerNames * []string, playerNumbers * []int, playerGoalNumbers * [] int){
	var playerName string
	var playerNumber int
	var playerGoalNumber int

	fmt.Println("Please enter your player name: ")
	fmt.Scan(&playerName)
	*playerNames = append(*playerNames, playerName)

	fmt.Println("Please enter your player number: ")
	fmt.Scan(&playerNumber)
	*playerNumbers = append(*playerNumbers, playerNumber)

	fmt.Println("Please enter your player goal total: ")
	fmt.Scan(&playerGoalNumber)
	*playerGoalNumbers = append(*playerGoalNumbers, playerGoalNumber)		
}

func removePlayer(removePlayerNames * []string){

	var removePlayerName string
	fmt.Scan(&removePlayerName)

	for i := 0; i < len(*removePlayerNames); i++{ 
		if removePlayerName == (*removePlayerNames)[i] {
			
		}
	}
}

func main() {
	for {
		var playerNames []string
		var playerNumbers []int
		var playerGoalNumbers []int

		showMenu()

		var choice int

		fmt.Println("Please enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addPlayer(&playerNames,&playerNumbers,&playerGoalNumbers)
			fmt.Println("Player name:", playerNames, "| Player number:", playerNumbers, "| Player Goal Total:", playerGoalNumbers)
			break
		case 2:
			
		case 3:

		case 4:

			fmt.Println("Most goals...")
		case 5:
			fmt.Println("Exiting...")
			break
		default:
			fmt.Println("Invalid choice. Please try again.")
			
		}
	}

}