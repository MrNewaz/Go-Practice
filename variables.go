package main

import "fmt"

func main() {
	var name string = "Saif"
	var age int = 124
	var salary float64 = 1000.50
	var isMarried bool = true

	dynamicType := "Dynamic Value"

	fmt.Println("Hello, my name is", name, "and I am", age, "years old. I earn", salary, "per year. I am married:", isMarried,
		"and my dynamic type is:", dynamicType)

	// Formatted String
	fmt.Printf("Hello, my name is %v and I am %v years old. I earn %v per year. I am married: %v and my dynamic type is: %v", name, age, salary, isMarried, dynamicType)
}
