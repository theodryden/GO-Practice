package main

import "fmt"

/* 
	Create a program asking for student data - name,age,grade

	Program has a menu with 3 options

	- Add student
	- List all Students
	- Exit

	// Use Pointers!! 
	*/


func main() {

	var names = []string{}
	var ages = []int{}
	var grades = []int{}

	var count int = 0
	var running bool = true
	var choice int

	fmt.Println(ages)

	for running == true {

		fmt.Println("1 - Add Student")
		fmt.Println("2 - List All Students")
		fmt.Println("3 - Exit")

		fmt.Println()
		fmt.Println("Please choose your option:")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Please enter the student's name:")
			fmt.Scan(&names[count])

			fmt.Println("Please enter the student's age:")
			fmt.Scan(&ages[count])

			fmt.Println("Please enter the student's grade:")
			fmt.Scan(&grades[count])
			
			count += 1

		} else if choice == 2 {
			for i:=0; i < count; i++ {
				fmt.Println("Name:", names[i])
				fmt.Println("Age:", ages[i])
				fmt.Println("Grade:", grades[i])
				fmt.Println()
			}
		} else if choice == 3 {
			fmt.Println("You have chosen to exit the system!")
			running = false
		} else {
			fmt.Println("Debug")
		}

	}



}