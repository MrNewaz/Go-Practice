package main

import "fmt"

func main() {
	var name string
	fmt.Printf("What is your name? : ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, Welcome to the game!\n", name)

	var age uint
	fmt.Printf("Enter your age? : ")
	fmt.Scan(&age)

	fmt.Println(age >= 18)
}
