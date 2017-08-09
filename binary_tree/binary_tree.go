package binaryTree

type Node struct {
  value int
  left *Node
  right *Node
}

func (node *Node) Insert(value int) {
  if value == node.value {
    return
  } else if value < node.value {
    if node.left == nil {
      node.left = &Node{ value: value }
    } else {
      node.left.Insert(value)
    }
  } else if value > node.value {
    if node.right == nil {
      node.right = &Node{ value: value }
    } else {
      node.right.Insert(value)
    }
  }
}

func (node *Node) Member(value int) bool {
  if value < node.value {
    if node.left != nil {
      return node.left.Member(value)
    }
  } else if value > node.value {
    if node.right != nil {
      return node.right.Member(value)
    }
  } else {
    return true
  }
  return false
}

func (node *Node) Delete(value int) *Node {
  if value < node.value {
    if node.left != nil {
      node.left = node.left.Delete(value)
    }
  } else if value > node.value {
    if node.right != nil {
      node.right = node.right.Delete(value)
    }
  } else {
    if node.isLeaf() {
      return nil
    } else if node.left == nil {
      return node.right
    } else if node.right == nil {
      return node.left
    } else {
      val, r_ptr := node.right.deleteMin()
      return &Node{ value: val, left: node.left, right: r_ptr }
    }
  }
  return node
}

func (node *Node) isLeaf() bool {
  return node.left == nil && node.right == nil
}

func (node *Node) deleteMin() (int, *Node) {
  if node.left == nil {
    return node.value, nil
  } else {
    v, l_ptr := node.left.deleteMin()
    node.left = l_ptr
    return v, l_ptr
  }
}
