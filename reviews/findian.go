package main

import (
  "fmt"
  "strings"
)

func main()  {
  var str string
  fmt.Printf("Enter String to find :")
  fmt.Scanln(&str)
  str1 := strings.ToUpper(str)

  if strings.HasPrefix(str1, "I") &&
      strings.HasSuffix(str1, "N") &&
      strings.Contains(str1, "A") {
        fmt.Println("Found!")
  } else {
      fmt.Println("Not Found!")
  }
}
