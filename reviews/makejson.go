package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What's your name? ")
	b, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	name := strings.TrimSpace(string(b))

	fmt.Print("What's your address? ")
	b, _, err = reader.ReadLine()
	if err != nil {
		panic(err)
	}

	address := strings.TrimSpace(string(b))

	user := map[string]string{
		"name":    name,
		"address": address,
	}

	b, err = json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
