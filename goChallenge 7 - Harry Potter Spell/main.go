package main

import (
	"fmt"
	s "strings"
)

func main() {

	// name of the spell 
	spell := "Avarda Kadavra"

	// use the count method from strings 
	fmt.Println("The character A is shown this many times:",s.Count(spell, "a"))
	
}


// https://www.tutorialspoint.com/how-to-count-the-number-of-repeated-characters-in-a-golang-string#:~:text=How%20to%20count%20the%20number%20of%20repeated%20characters%20in%20a%20Golang%20String%3F,-Go%20ProgrammingServer&text=Count()%20is%20a%20built,character%2Fstring%20in%20a%20string.