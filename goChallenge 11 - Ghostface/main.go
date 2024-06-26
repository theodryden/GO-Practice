package main

import (
	"fmt" 
	"strings"
)

//Homework
func main() {
	// Call Logs
	callLogs := []string{
		"Ghostface: Hello, Sidney",
		"Mom: Dinner's ready",
		"Ghostface: Do you like scary movies?",
		"Friend: Are you calling me",
	}

	// Looking for the words in the array 
	caller := "Ghostface"

	// Initialize count - anytime using a counter
	count := 0

	// Iterate over the call logs and ocunt the occurences of the caller
	for _,log := range callLogs {

		// Checks if the log starts with the caller's name
		if strings.HasPrefix(log,caller+":"){
			count++
		}
		
	}
	// Print the results
	fmt.Printf("The caller %s appears %d times in the call logs.\n",caller,count)

}