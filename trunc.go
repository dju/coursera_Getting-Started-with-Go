// https://www.coursera.org/learn/golang-getting-started/
// Week 2 : Peer-graded Assignment: Module 2 Activity: trunc.go
/* Instructions
The goal of this assignment is to practice working with floating point numbers and truncation code in Go.
Review criteria
This assignment is worth a total of 10 points:
	3 points will be given if a program is written.
	2 points will be given if the program compiles correctly.
	3 points will be given for the first test execution,
		if the program correctly prints the truncated integer when a floating point number is entered.
	2 points will be given for the second test execution,
		if the program correctly prints the truncated integer when another floating point number is entered.
*/

package main

import "fmt"

func main() {
	var response float32
	var err error
	fmt.Printf("Enter a float: ")
	_, err = fmt.Scan(&response)
	if err != nil {
		fmt.Printf("Error -->%s<-*\n", err)
	} else {
		fmt.Printf("You have entered : %f\n", response)
		fmt.Printf("Truncated integer: %d\n", int(response))
	}
}
