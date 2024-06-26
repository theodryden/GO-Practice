package main

import "fmt"

func main() {
	
	//Slices
	var names []string
	var brands []string
	var quantities []int

	for {
	// Menu
	fmt.Println("Here's a menu of options, select an option!")
	fmt.Println("1 - Add new item:")
	fmt.Println("2 - Remove an item:")
	fmt.Println("3 - Display all items:")
	fmt.Println("4 - Exit the program:")

	// Select an item from the menu
	choice := 0 
	fmt.Scan(&choice)

	// Run the corresponding functionality : Add/Remove/Update
	var newItem string
	var newItemBrandName string
	var newItemQuantity int
	var elementToRemove string
	
	// 1 - Add Item Block
	if choice == 1 {
		fmt.Println("1. You've chosen to add a new item.\nPlease enter it's name:")
		fmt.Scan(&newItem)

		fmt.Println("2. Now enter the brand:")
		fmt.Scan(&newItemBrandName)

		fmt.Println("3. Now enter the quantity:")
		fmt.Scan(&newItemQuantity)

		// Add new item, brand, and quantity to the slice
		for {
			names = append(names, newItem)
			brands = append(brands, newItemBrandName)
			quantities = append(quantities, newItemQuantity)
			break
		}

	// 2- Remove Item Block
	} else if choice == 2 {
		
		// Get the elemet to remove string slice
		fmt.Println("You've chosen to remove an new item.\nPlease enter it's name to remove:")
		fmt.Scan(&elementToRemove)
		
		// Find the index of the element to remove
		indexToRemove := -1

		for i, name := range names {
			if name == elementToRemove {
				indexToRemove = i
				break
			}
		}

		// If the element is found, remove it
		if indexToRemove != -1 {

			names = append(names[:indexToRemove], names[indexToRemove+1:]...)

			fmt.Println("This has now been removed. Remaining names:", names)

		} else {
			fmt.Println("Element not found in the names slice.")
		}

	// Display All Items
	} else if choice == 3 {

		// Loop through the names in the names array
		fmt.Println("You've chosen to display all items!")

		for i := 0; i < len(names); i++ {
			fmt.Println("Item:",names[i])
		}

		for i := 0; i < len(brands); i++ {
			fmt.Println("Brand:",brands[i])
		}

		for i := 0; i < len(quantities); i++ {
			fmt.Println("Quantity:",quantities[i])
		}
	} else {
		fmt.Println("You've chosen to exit this program.\nSee you again soon!")
		break
	}
	
}

	
}