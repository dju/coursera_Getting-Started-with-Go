package main

import (
  "fmt"
  "sort"
  "bufio"
  "os"
  "strconv"
)

func main() {
   var numbers = make([]int, 0, 3)
   fmt.Println("Enter the Number (X to Exit) :")
   scanner := bufio.NewScanner(os.Stdin)

   for scanner.Scan() {
     var strVal string = scanner.Text()
     if strVal == "X" {
       fmt.Println("Thanks for adding Numbers!")
       os.Exit(0)
    }
    intVal, err := strconv.Atoi(strVal)
    if err != nil {
  		fmt.Println(err)
  		return
  	}
      numbers = append(numbers, intVal)

      sort.Ints(numbers)
      fmt.Println(numbers)
   }
}
