/* read.go */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname []rune // 20 character limit is enforced during assignment.
	lname []rune
}

func main() {
	
	var fileName string
	var names []Name

	fmt.Printf("Enter the data filename: ")
	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	
	if (err != nil) {
		println("Could not read data file, please check the filename/path and try again.")
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for (scanner.Scan()) {
		n := new(Name)

		line := strings.Fields(scanner.Text())
		l0 := []rune(line[0])
		l1 := []rune(line[1])

		// Enforce 20 character limit.
		if (len(l0) > 20) {
			l0 = l0[:20]
		}

		if (len(l1) > 20) {
			l1 = l1[:20]
		}

		n.fname = []rune(l0)
		n.lname = []rune(l1)
		names = append(names, *n)
	}

	for _, v := range names {
		fname := string(v.fname[:])
		lname := string(v.lname[:])
		println(fname + " " + lname)
	}
}