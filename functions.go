package main

import "fmt"

func add(x int, y int) (z int) {
	z = x + y
	return

}

func functions() {
	x := add

	fmt.Println(x(5, 10))

	test := func(x int, y int) int {
		return x + y
	}(5, 10)

	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)
	fmt.Println(test)
}
