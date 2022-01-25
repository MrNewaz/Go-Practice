package main

import "fmt"

type Point struct {
	x int
	y int
	// isOnGrid bool
}

func main() {
	p1 := Point{x: 2}
	p2 := Point{y: 3}
	fmt.Println("p1: ", p1, "\np2: ", p2)
}
