package main

import "fmt"

func main() {
	var name string
	fmt.Printf("What is your name?: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, Welcome to the game!\n", name)

	var age uint
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Printf("Sorry, you are not eligible to play this game.\n")
		return
	} else {
		fmt.Printf("You are eligible to play this game.\n\n")

	}

	var score int = 0
	fmt.Println("Current Score: ", score)

	fmt.Printf("Which is better, \n1.RTX 3080 \nor \n2.GTX 950?\n(1/2):)")
	var gpu string
	fmt.Scan(&gpu)

	if gpu == "1" {
		fmt.Printf("You have chosen the best GPU ever!\n")
		score++
	} else {
		fmt.Printf("You have chosen the worst GPU ever!\n")
	}

	fmt.Println("Current Score: ", score)

	fmt.Printf("Which is better, \n1.Intel Core i7 \nor \n2.AMD Ryzen?\n(1/2): ")
	var cpu string
	fmt.Scan(&cpu)
	if cpu == "2" {
		fmt.Printf("You have chosen the best CPU ever!\n")
		score++
	} else {
		fmt.Printf("You have chosen the worst CPU ever!\n")
	}

	fmt.Println("Game Over, Your Final Score: ", score)
}
