package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go World")
	name := "Prosper"
	var email string = "prospermbuma@gmail.com"
	var age int = 30
	greetings(name)

	fmt.Println()
	fmt.Println("======== User Profile ========")
	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
	fmt.Println("Age:", age)
	fmt.Println("==============================")

	fmt.Println()
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numbers {
		fmt.Println(num)
	}

	fmt.Println()
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}
}

func greetings(name string) {
	fmt.Println("Hello " + name)
}
