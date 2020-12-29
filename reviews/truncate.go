package main

import "fmt"

func main() {
	fmt.Println("Truncate floating numbers into integers. (ctrl-c to exit)")
	for {
		var input float32
		fmt.Print("input a floating number: ")
		fmt.Scan(&input) // ignore return values
		output := int(input)
		fmt.Printf("truncated number = %d\n", output)
	}
}
