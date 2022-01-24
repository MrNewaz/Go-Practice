package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type Something?: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed, %v", input)

}
