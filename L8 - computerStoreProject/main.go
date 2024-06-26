package main

import (
	"fmt"
)

// Ask the user for computer date (name,price,os) The program should have 5 options:
/*
- Add computer
- Remvoe computer
- Edit computer by name
- List all computers
- Exit
*/

/* func addComputer(names []string, prices []float32, os []string){


} */

func listComputers(names []string, prices []float32, os []string){
	for i := 0; i < len(names); i++ {
		fmt.Println("Name:", names[i], "| Price:", prices[i],"| Operating System:", os[i])
	}
}

func main() {
	
	var names = []string{}
	var prices = []float32{}
	var os = []string{}

	var running bool = true
	var choice int

	for running == true {

		fmt.Println("1 - Add Computer")
		fmt.Println("2 - Remove Computer")
		fmt.Println("3 - Edit Computer By Name")
		fmt.Println("4 - List All Computers")
		fmt.Println("5 - Exit")

		fmt.Println("Please enter your choice:")
		fmt.Scan(&choice)

		if choice == 1 {

			var name string
			var price float32
			var operatingSystem string
		
			fmt.Println("Please enter the name:")
			fmt.Scan(&name)
		
			fmt.Println("Please enter the price:")
			fmt.Scan(&price)
		
			fmt.Println("Please enter the os:")
			fmt.Scan(&operatingSystem)
		
			names = append(names, name)
			prices = append(prices, price)
			os = append(os, operatingSystem)
			
		} else if choice == 2 {
			
		} else if choice == 3 {
			var name string
			fmt.Println("Please enter the name:")
			fmt.Scan(&name)

			// Edit computer details by name
			for i := 0; i < len(names); i++ {

				if names[i] == name {
					var price float32
					var operatingSystem string

					fmt.Println("Please enter the new price")
					fmt.Scan(&price)

					fmt.Println("Please enter the new os")
					fmt.Scan(&operatingSystem)

					prices[i] = price
					os[i] = operatingSystem

				}
			}


		} else if choice == 4 {
			// List all computer information
			listComputers(names,prices,os)

		} else if choice == 5 {
			running = false

		}


	}



}