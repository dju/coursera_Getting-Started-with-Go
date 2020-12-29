// https://www.coursera.org/learn/golang-getting-started/
// Week 2 : Peer-graded Assignment: Module 2 Activity: findian.go
/* Instructions
The goal of this assignment is to practice working with strings in Go.
Review criteria
This assignment is worth a total of 10 points:
3 points will be given if a program is written.
2 points will be given if the program compiles correctly.
3 points will be given for the first test execution,
	if the program correctly prints "Found!"
	when a string that contains the characters ‘i’, ‘a’, and ‘n’, such as "iaaaan" is entered.
2 points will be given for the second test execution,
	if the program correctly prints "Not Found!"
	when a string that does not contain the characters ‘i’, ‘a’, and ‘n’ is entered.
 */
/*
Examples: The program should print “Found!”
	for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!”
	for the following strings, “ihhhhhn”, “ina”, “xian”.
 */
package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
func main() {
	var err error
	var response, trim_response string
	var first_i, second_a, third_n int

	fmt.Printf("Enter your string: ")
	in := bufio.NewReader(os.Stdin)
	response, err = in.ReadString('\n')
	// _, err = fmt.Scan(&response)
	if err != nil {
		fmt.Printf("Error -->%s<-*\n",err)
	} else {
		trim_response = strings.ToLower(strings.TrimSpace(response))
		first_i  = strings.IndexByte(trim_response, 'i')
		second_a = strings.IndexByte(trim_response, 'a')
		third_n  = strings.IndexByte(trim_response, 'n')
		//fmt.Printf("response -->%s<--\ntrim_response -->%s<--\nfirst_i %d,second_a %d,third_n %d\n",response, trim_response, first_i, second_a, third_n)
		if (first_i == 0 ) && (first_i < second_a) && ( second_a < third_n) {
			fmt.Printf("Found!\n")
		} else {
			fmt.Printf("Not Found!\n")
		}
	}
}
