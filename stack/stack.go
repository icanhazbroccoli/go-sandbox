package stack

type Stack struct {
  data [10]int
  i int
}

func (s *Stack) Push(el int) {
  s.data[s.i] = el
  s.i++
}

func (s *Stack) Pop() int {
  res := s.data[s.i - 1]
  s.i--
  s.data[s.i] = 0
  return res
}
