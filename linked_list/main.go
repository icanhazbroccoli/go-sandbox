package main

import (
  "fmt"
  "container/list"
)

type NodeElement interface {}

type Node struct {
  value NodeElement
  left  *Node
  right *Node
}

func (node *Node) PushBack(element NodeElement) {
  newNode := Node{ value: element }
  node.right = &newNode
  newNode.left = node
}

func main() {
  list := new(list.List)
  list.Init()
  list.PushBack(1)
  list.PushBack(2)
  list.PushBack(4)
  fmt.Println(list)

  myList := Node{value: 1}
  myList.PushBack(2)
  myList.PushBack(4)
  fmt.Println(myList)
}
