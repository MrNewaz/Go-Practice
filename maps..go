package main

import "fmt"

func maps() {
	var mp map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println(mp)

	// Another Way

	mp2 := make(map[string]int)

	mp2["a"] = 1
	mp2["b"] = 2
	mp2["c"] = 3

	fmt.Println(mp2)
}
