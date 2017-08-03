package main

import "fmt"

func bubble_sort(a []int) ([]int) {
  for {
    swapped := false
    for i := 1; i <= len(a) - 1; i++ {
      if a[i - 1] > a[i] {
        tmp := a[i]
        a[i] = a[i - 1]
        a[i - 1] = tmp
        swapped = true
      }
    }
    if !swapped {
      break
    }
  }
  return a
}

func main() {
  arr := []int{ 1, 0, 15, 0, 12, 42 }
  fmt.Println(bubble_sort(arr))
}
