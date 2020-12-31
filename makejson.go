// https://www.coursera.org/learn/golang-getting-started/
// Week 4 : Peer-graded Assignment: Module 4 Activity: makejson.go
/* Instructions
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
 */
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)
func main() {
	m := map[string]string {"name":"", "address":""}
	in := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter name :\n")
	in.Scan()
	m["name"] = in.Text()
	fmt.Printf("Enter address :\n")
	in.Scan()
	m["address"] = in.Text()

	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(json))
//	os.Stdout.Write(json)
}
