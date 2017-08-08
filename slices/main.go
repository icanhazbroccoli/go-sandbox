package main

import "fmt"

func main() {
  a := [...]int{ 1, 2, 3, 4, 5, }
  fmt.Println(a[2:4]);
  fmt.Println(a[1:5]);
  fmt.Println(a[:]);
  fmt.Println(a[2:4:5])
}
