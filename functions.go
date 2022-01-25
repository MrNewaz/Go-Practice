package main

import "fmt"

func add(x int, y int) (z int) {
	z = x + y
	return

}

func main() {
	x := add(10, 5)
	fmt.Println(x)
}
