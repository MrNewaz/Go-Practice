package main

import "fmt"

func main() {
	var name string
	fmt.Printf("What is your name? ")

	fmt.Scan(&name)

	fmt.Println("Hello, my name is", name)

}
