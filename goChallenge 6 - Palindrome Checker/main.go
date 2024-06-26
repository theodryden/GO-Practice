package main

import "fmt"

func isPalindrome(word string) bool{

	// while the index is less than letters in the world split in half, increment it
	for i := 0; i < len(word)/2; i++ {
		// if the start of the word's individual letters are not equal to end of the word, return false
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	// return true boolean
	return true
}

func main() {

	// Test words using Palindrome function
	fmt.Println(isPalindrome("deed"))
	fmt.Println(isPalindrome("peep"))
	fmt.Println(isPalindrome("reconnaisance"))
	
}




/*package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("anna"))
	fmt.Println(isPalindrome("not a palindrome"))
}
*/

//https://go.dev/play/p/5FUOzjBa-o