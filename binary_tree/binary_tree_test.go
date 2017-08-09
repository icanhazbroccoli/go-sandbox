package binaryTree

import (
  "testing"
)

func TestMember(t *testing.T) {
  tree := Node{
    value: 1,
    left:  &Node{ value: 0 },
    right: &Node{ value: 2 },
  }
  if tree.Member(1) != true {
    t.Log("1 is not a member")
    t.Fail()
  }
  if tree.Member(0) != true {
    t.Log("0 is not a member")
    t.Fail()
  }
  if tree.Member(2) != true {
    t.Log("2 is not a member")
    t.Fail()
  }
}

func TestInsert(t *testing.T) {
  tree := Node{ value: 10 }
  tree.Insert(0)
  tree.Insert(20)
  if tree.Member(0) != true {
    t.Log("0 is not a member")
    t.Fail()
  }
  if tree.Member(20) != true {
    t.Log("20 is not a member")
    t.Fail()
  }
  if tree.left == nil {
    t.Log("left child is nil")
    t.Fail()
  }
  if tree.right == nil {
    t.Log("Right child is nil")
    t.Fail()
  }
  if tree.left.value != 0 {
    t.Log("Left value != 0")
    t.Fail()
  }
  if tree.right.value != 20 {
    t.Log("Right value != 20")
    t.Fail()
  }
}

func TestDelete1(t *testing.T) {
  tree := Node{
    value: 10, left: &Node{ value: 0 }, right: &Node{ value: 20 },
  }
  tree = *tree.Delete(0)
  if tree.Member(0) {
    t.Log("0 is still a member")
    t.Fail()
  }
  if tree.Member(20) != true {
    t.Log("20 is not a member")
    t.Fail()
  }
  if tree.Member(10) != true {
    t.Log("10 is not a member")
    t.Fail()
  }
  if tree.left != nil {
    t.Log("Left child is not nil")
    t.Fail()
  }
}

func TestDelete2(t *testing.T) {
  tree := Node{
    value: 10, left: &Node{ value: 0 }, right: &Node{ value: 20 },
  }
  tree = *tree.Delete(10)
  if tree.Member(10) {
    t.Log("10 is still a member")
    t.Fail()
  }
  if tree.Member(20) != true {
    t.Log("20 is not a member")
    t.Fail()
  }
  if tree.Member(0) != true {
    t.Log("0 is not a member")
    t.Fail()
  }
  if tree.right != nil {
    t.Log("Right child is not nil")
    t.Fail()
  }
}
