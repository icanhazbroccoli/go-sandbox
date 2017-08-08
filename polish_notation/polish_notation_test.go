package polishNotation

import "testing"

func TestAdd(t *testing.T) {
  if Calculate("1 1 +") != 2 {
    t.Log("1 + 1 != 2")
    t.Fail()
  }
}

func TestSubstr(t *testing.T) {
  if Calculate("3 2 -") != 1 {
    t.Log("3 - 2 != 1")
    t.Fail()
  }
}

func TestMult(t *testing.T) {
  if Calculate("2 2 *") != 4 {
    t.Log("2 * 2 != 4")
    t.Fail()
  }
}

func TestDiv(t *testing.T) {
  if Calculate("5 2 /") != 2 {
    t.Log("int(5 / 2) != 2")
    t.Fail()
  }
}

func TestComplex(t *testing.T) {
  if Calculate("2 2 + 5 - 10 *") != -10 {
    t.Log("((2 + 2) - 5) * 10 != -10")
    t.Fail()
  }
}
