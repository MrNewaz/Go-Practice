package main

import "fmt"

func main() {
	fmt.Printf("Hello, World %T %v", 10, 20)
	var x string = fmt.Sprintf("Hege %v", 20)
	fmt.Println(x)
}