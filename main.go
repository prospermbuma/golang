package main

import (
	"fmt"
)

// GLOBAL VARIABLES
var name string = "Prosper"
var email string = "prospermbuma@gmail.com"
var age int = 30

// MAIN FUNCTION
func main() {
	fmt.Println("Hello Go World")
	greetings(name)

	fmt.Println()
	fmt.Println("======== User Profile ========")
	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
	fmt.Println("Age:", age)
	fmt.Println("==============================")

	fmt.Println()
	// ARRAY - Is the list with a fixed size, for example 10 numbers.
	// Loop through an array
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numbers {
		if num%2 == 0 && num > 2 && num < 10 {
			fmt.Println(num)
		}
	}

	fmt.Println()
	// SLICE - Is the list with NO fixed size
	// Loop through a slice
	names := []string{"Victoria", "Prosper", "Joseph", "Gladness"}
	for index, name := range names {
		index++
		//fmt.Print(index, ". ", name, "\n")
		fmt.Printf("%d. %s\n", index, name)
	}

	fmt.Println()
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}

// FUNCTION DECLARATION
// User-defined functions
func greetings(name string) {
	fmt.Println("Hello " + name)
}
