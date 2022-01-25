package main

import "fmt"

func slices() {
	var s []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := append(s, 11)
	fmt.Println(s, "\n", b)
}
