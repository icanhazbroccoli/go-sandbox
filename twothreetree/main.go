package main

type Node interface {
	IsLeaf() bool
	insert1(element int) (Node, int)
}

type Node23 struct {
	firstchild  *Node
	secondchild *Node
	thirdchild  *Node
	lowofsecond int
	lowofthird  int
}

func (node Node23) IsLeaf() bool {
	return false
}

func (node Node23) insert1(element int) (pnew Node, low int) {

}

type Leaf struct {
	element int
}

func (leaf Leaf) IsLeaf() bool {
	return true
}

func (leaf Leaf) insert1(element int) (pnew Node, low int) {
	if element != leaf.element {
		pnew = Leaf{element: element}
		low = element
	}
	return pnew, low
}

func main() {

}
