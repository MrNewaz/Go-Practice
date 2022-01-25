package main

import "fmt"

func ranges() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, element := range a {
		fmt.Println(i, element)
	}

	//  Anonymous variable "_"
	for _, element := range a {
		fmt.Println(element)
	}

}
