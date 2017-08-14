package main

import "fmt"

type Person struct {
  name string
  age int
}

func main() {
  var p1 Person
  p2 := new(Person)

  fmt.Println(p1)
  fmt.Println(p2)
}
