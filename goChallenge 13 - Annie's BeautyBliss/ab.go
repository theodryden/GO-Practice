package main

import "fmt"

func main() {

	// Run the program

	clientName := []string{}
	time := []string{}
	duration := []int{}

	var cNameSlice string
	var timeSlice string
	var durationSlice int

	for {
	// Menu
	fmt.Println("Welcome to your menu, here are your options:")
	fmt.Println("1 - Add an appointment.")
	fmt.Println("2 - Cancel an appointment.")
	fmt.Println("3 - List all appointments.")
	fmt.Println("4 - Exit the program.")

	// Take User Input
	var userChoice int
	fmt.Println("Please select an option:")
	fmt.Scan(&userChoice)

	// Add an appointment functionality
	if userChoice == 1 {

		fmt.Println("Enter your client name:")
		fmt.Scan(&cNameSlice)

		fmt.Println("Enter your appointment date:")
		fmt.Scan(&timeSlice)

		fmt.Println("Enter your appointment duration:")
		fmt.Scan(&durationSlice)
	
		clientName = append(clientName, cNameSlice)
		time = append(time, timeSlice)
		duration = append(duration, durationSlice)
		
	// Remove the entry	
	} else if userChoice == 2 {
		var removeName string
		fmt.Println("You have chosen to cancel an appointment!")
		fmt.Println("Enter name to cancel:")
		fmt.Scan(&removeName)

		removeIndex := -1

		// Step 1 Find index of element to remove in client name 
		for i := 0; i < len(clientName); i++ {
			if removeName == clientName[i] {
				removeIndex = i
			}
		}

		// Step 2 - update element index in client name to remove it
		for i := 0; i < len(clientName); i++ {
			
			fmt.Println("You have chosen to cancel an appointment:")

			if i == removeIndex { 

				clientName = append(clientName[:removeIndex], clientName[removeIndex+1:]...)
				fmt.Printf("Client %s's appointment has now been removed.\n", removeName)
				
			} else {
				fmt.Printf("error")
			}
		}

	// Show all the appointments
	} else if userChoice == 3 {

		for i := 0; i < len(clientName); i++ {

			fmt.Printf("Name: %s, Appointment Date: %s, Duration: %v\n",clientName[i],time[i],duration[i])
			
		}


	} else if userChoice == 4 {
		fmt.Println("You have chosen to exit the program:")
		break
	}

}




}