package polishNotation

import (
  "fmt"
  "go-sandbox/stack"
  "strings"
  "strconv"
)

func add(s *stack.Stack) {
  a := s.Pop()
  b := s.Pop()
  s.Push(a + b)
}

func substr(s *stack.Stack) {
  b := s.Pop()
  a := s.Pop()
  s.Push(a - b)
}

func mult(s *stack.Stack) {
  a := s.Pop()
  b := s.Pop()
  s.Push(a * b)
}

func div(s *stack.Stack) {
  b := s.Pop()
  a := s.Pop()
  s.Push(a / b)
}

func Calculate(s string) int {
  var stack stack.Stack
  chunks := strings.Split(s, " ")
  for _, sym := range(chunks) {
    if sym == "+" {
      add(&stack)
    } else if sym == "-" {
      substr(&stack)
    } else if sym == "*" {
      mult(&stack)
    } else if sym == "/" {
      div(&stack)
    } else {
      i, err := strconv.Atoi(sym)
      if err != nil {
        panic(fmt.Sprintf("Unable to parse symbol %s\n", sym))
      }
      stack.Push(i)
    }
  }
  return stack.Pop()
}
