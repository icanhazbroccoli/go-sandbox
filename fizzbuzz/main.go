package main

import (
  "fmt"
  "strconv"
)

func fizzbuzz(i int) (fizz, buzz string) {
  if i % 3 == 0 { fizz = "Fizz" }
  if i % 5 == 0 { buzz = "Buzz" }
  if !( len(fizz) > 0 || len(buzz) > 0 ) {
    fizz = strconv.Itoa(i)
  }
  return
}

func main() {
  for i := 1; i <= 100; i++ {
    fizz, buzz := fizzbuzz(i)
    fmt.Print(fizz, buzz, "\n")
  }
}
