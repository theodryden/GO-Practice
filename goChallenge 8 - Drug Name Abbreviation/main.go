package main

import (
	"fmt"
	s "strings"
)


func main() {
	// strings.Fields() method
	var drug string = "Acetaminophen Codeine Phosphate"
	//fmt.Println("The drug abbreviated is:",drug)

	shortHand := s.Fields(drug)

	fmt.Println(shortHand)

	// loop through the words in shortHand variable
	for _, word := range shortHand {
		fmt.Print(string(word[0])) //print the first letter of each word
	 }

}

// https://www.tutorialspoint.com/golang-program-to-print-first-letter-of-each-word-using-regex#:~:text=Field()%20function,each%20word%20inside%20the%20loop.