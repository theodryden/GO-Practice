package main

import (
	"fmt"
	"strings"
)

func main() {

	var flower string
	var vowels string = "aeiouAEIOU"
	// for exiting the loop
	var exit bool = false
	var count int

	for exit == false {

		fmt.Println("Please enter the name of a flower: ")
		fmt.Scan(&flower)

		if flower == "exit" {
			exit = true 
			continue
		}

		count = 0 

		// For loop version 1 you can use
		for _, letter := range flower {
			if strings.ContainsRune(vowels, letter) {
				count += 1
			}
			
		} 

		fmt.Println("The number of vowels is: ", count)

		// For Loop version 2 you can use - Pick one or the other
		runes := []rune(flower)

		for i := 0; i < len(flower); i ++ {
			if strings.ContainsRune(vowels, runes[i]) {
				count += 1
			} 
		}
		fmt.Println("The number of letters is: ", count)
		
	}
	
	
}