package main

import (
	"fmt"
)

func main() {

	var f float64

	fmt.Println("Please enter a floating point number:")
	_,err := fmt.Scanf("%f", &f)
	if err != nil {
		fmt.Println("Please enter a floating point number")
	} else {
		fmt.Printf("You entered: %f\n", f)
		fmt.Printf("The integer is: %d", int(f))
	}
}
