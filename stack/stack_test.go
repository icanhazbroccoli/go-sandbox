package stack

import "testing"

func TestStack(t *testing.T) {
  var s Stack
  s.Push(1)
  s.Push(2)
  s.Push(3)
  if s.Pop() != 3 {
    t.Log("The last element should be 3")
    t.Fail()
  }
}
