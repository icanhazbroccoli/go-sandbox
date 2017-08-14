package main

import "fmt"

type enumerable interface {}

func Map(collection []enumerable, cb func(enumerable)enumerable) []enumerable {
  res := make([]enumerable, len(collection))
  for k, v := range collection {
    res[k] = cb(v)
  }
  return res
}

func doubleOrSquare(el enumerable) enumerable {
  switch el.(type) {
    case int:
      return el.(int) * el.(int)
    case string:
      return el.(string) + el.(string)
    default:
      return el
  }
}

func main() {
  ints := []enumerable{1,2,3,4,5,6,7}
  strings := []enumerable{ "a", "b", "c", "d", "e" }
  fmt.Println(Map(ints, doubleOrSquare))
  fmt.Println(Map(strings, doubleOrSquare))
}
