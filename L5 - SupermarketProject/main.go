package main

import "fmt"

// Pointers are useful when you're changing the data
// Pointers are not needed when you are not changing the data i.e Appending a slice

/*
  Create a program that asks the user for unlimited product data (name, price, stock)
  That program should have a menu with this options:

  1 - Add product
  2 - Show product by name
  3 - Remove product
  4 - Edit Product
  5 - List all products
  6 - Exit

*/

func addProduct(pointerNames*[] string, pointerPrices*[] float32, pointerStocks*[] int){
	fmt.Println("Enter the name of the new product:")
	var name string
	fmt.Scan(&name)

	var price float32

	for price <= 0 {
		fmt.Println("Enter the price of the new product:")
		fmt.Scan(&price)
	}

	var stock int = -1

	for stock <= 0 {
		fmt.Println("Enter the stock of the new product:")
		fmt.Scan(&stock)
	}
			
	*pointerNames = append(*pointerNames, name)
	*pointerPrices = append(*pointerPrices, price)
	*pointerStocks = append(*pointerStocks, stock)
}

func showProduct(names[] string, prices[] float32, stocks[] int){
	var name string
	fmt.Println("Please enter the name of the product you want to list:")
	fmt.Scan(&name)

	for i:=0; i < len(names); i++ {
		if names[i] == name {
			fmt.Println("Name:", names[i], "|", "Price:", prices[i], "|", "Stock:", stocks[i])
			break
		} 
	}

}

/*
Always use (*pointerNames)[i] when pointerNames is a pointer to a slice in an if statement, 
and you need to access an element of the slice it points to. 
The brackets ensure the correct order of operations: dereferencing first, then indexing.
*/
func removeProduct(pointerNames*[] string, pointerPrices*[] float32, pointerStocks*[] int){
	var removeName string
	fmt.Println("Enter name to remove")
	fmt.Scan(removeName) 

	for i := 0; i < len(*pointerNames); i++ {
		if (*pointerNames)[i] == removeName {
			*pointerNames = append((*pointerNames)[:i], (*pointerNames)[i+1:]...)
			*pointerPrices = append((*pointerPrices)[:i], (*pointerPrices)[i+1:]...)
			*pointerStocks = append((*pointerStocks)[:i], (*pointerStocks)[i+1:]...)
			break
		}
		fmt.Printf("Names: %s| Prices: %v| Stocks: %v|", (*pointerNames)[i],(*pointerPrices)[i],(*pointerStocks)[i])
	}
}


func main() {

	names := []string{}
	prices :=[]float32{}
	stocks :=[]int{}

	choice := 0

	for true {
		fmt.Println("1 - Add Product")
		fmt.Println("2 - Show Product By Name")
		fmt.Println("3 - Remove Product")
		fmt.Println("4 - Edit Product")
		fmt.Println("5 - List All Products")
		fmt.Println("6 - Exit")

		fmt.Println()

		fmt.Println("Please enter an option:")
		fmt.Scan(&choice)

		if choice == 1 {

			addProduct(&names,&prices,&stocks)
/*
			fmt.Println("Enter the name of the new product:")
			var name string
			fmt.Scan(&name)


			var price float32

			for price <= 0 {
				fmt.Println("Enter the price of the new product:")
				fmt.Scan(&price)
			}

			var stock int = 0

			for stock <= 0 {
				fmt.Println("Enter the stock of the new product:")
				fmt.Scan(&stock)
			}
			
			names = append(names, name)
			prices = append(prices, price)
			stocks = append(stocks, stock) */

		} else if choice == 2 {

			showProduct(names,prices,stocks)

		} else if choice == 3 {

		/*	var removeName string
			fmt.Println("Enter name to remove")
			fmt.Scan(removeName) 

			for i := 0; i < len(names); i++ {
				if names[i] == removeName {
					names = append(names[:i], names[i+1:]...)
					prices = append(prices[:i], prices[i+1:]...)
					stocks = append(stocks[:i], stocks[i+1:]...)
					break
				}
				fmt.Printf("Names: %s| Prices: %v| Stocks: %v|", names[i],prices[i],stocks[i])
			
			} */

		} else if choice == 4 {
			// Re-Assign Product
			
				var name string

				fmt.Println("What product to change?:")
				fmt.Scan(&name)

				for i := 0; i < len(names); i++ {

					if name == names[i] {

						for prices[i] <= 0 {
						fmt.Println("What's the new price?:")
						fmt.Scan(&prices[i])
						}

						for stocks[i] <= 0 {
						fmt.Println("What's the new stock?:")
						fmt.Scan(&stocks[i])
						}

						break

					} else {
						fmt.Println("Not possible")
					
					}
				}
				
		} else if choice == 5 {

			for i:=0; i < len(names); i++ {
				fmt.Println("Name:", names[i], "|", "Price:", prices[i], "|", "Stock:", stocks[i])
			}

		} else if choice == 6 {
			break
		}

	} 

}
