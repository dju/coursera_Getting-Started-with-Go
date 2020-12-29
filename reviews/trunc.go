package main

import (
	"fmt"
)

func main() {
	var input float32
	fmt.Printf("Enter a float: ")
	fmt.Scan(&input)
	fmt.Printf("Truncated float: %d\n", int(input))
}
