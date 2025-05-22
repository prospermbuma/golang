package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go World")
	name := "Prosper"
	greetings(name)
	for i := 1; i <= 10; i++ {
		fmt.Println("i =", 100/i)
	}
}

func greetings(name string) {
	fmt.Println("Hello " + name)
}
