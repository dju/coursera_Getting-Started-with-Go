// https://www.coursera.org/learn/golang-getting-started/
// Week 3 : Peer-graded Assignment: Module 3 Activity: slice.go
/* Instructions
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3.
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	const size int = 3
	var err error
	var response string
	var currentValue int
	var quit bool
	var mySlice = make([]int, size)
	var tmpMySlice = make([]int, size)
	var currentPos int = 0

	for !quit {
		fmt.Printf("Enter integer (X to quit) :\n")
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		response = in.Text()
		if response == "X" || response == "x" {
			quit = true
		} else {
			if currentValue, err = strconv.Atoi(response); err == nil {
				if currentPos >= len(mySlice) {
					mySlice = append(mySlice, currentValue)
				} else {
					tmpMySlice[currentPos] = currentValue
					copy(mySlice,tmpMySlice)
				}
				sort.Ints(mySlice)
				currentPos++
			} else {
				fmt.Printf("Error -->%s<-*\n", err)
			}
		}
		fmt.Println("the slice in sorted order", mySlice)
	}
	fmt.Println("goodbye.")
}
