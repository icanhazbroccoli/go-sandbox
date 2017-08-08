package main

import "fmt"

func max(list []int) (int) {
  max_el := list[0]
  for _, el := range list[1:] {
    if el > max_el {
      max_el = el
    }
  }
  return max_el
}

func mmap(list []int, f func(int)(int)) ([]int) {
  ret := make([]int, len(list))
  for ix, el := range(list) {
    ret[ ix ] = f(el)
  }
  return ret
}

func main() {
  fmt.Println(mmap([]int{1, 2, 3, 4, 5}, func(i int) (int) { return i * i }))
}
