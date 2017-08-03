package main

import "fmt"

func main() {
  fmt.Println(fibonacci(10))
}

func fibonacci(n int) (sequence []int) {
  sequence = make([]int, n)
  copy(sequence, []int{1,1})
  for i := 2; i < n; i++ {
    sequence[i] = sequence[i - 2] + sequence[i - 1]
  }
  return
}
