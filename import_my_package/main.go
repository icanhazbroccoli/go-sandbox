package main

import (
  "fmt"
  "even"
)

func main() {
  i := 5
  fmt.Printf("%d is %v\n", i, even.Even(i))
}
