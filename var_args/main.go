package main

import "fmt"

func main() {
  print_numbers(1,2,3,4,5,6,7)
}

func print_numbers(args ...int) {
  for _, n := range args {
    fmt.Println(n)
  }
}
