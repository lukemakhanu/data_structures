package main

import "fmt"

// Node : binary search tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert : used to insert into the tree.
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			// key is empty so create a node here.
			n.Right = &Node{Key: k}
		} else {
			// there was a key already do another recursion but on the right child.
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			// key is empty so create a node here.
			n.Left = &Node{Key: k}
		} else {
			// there was a key already do another recursion but on the left child.
			n.Left.Insert(k)
		}
	}
}

// Search : used to search the tree.
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Key < k {
		// move to the right child
		return n.Right.Search(k)
	} else if n.Key > k {
		// move to the left child.
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tr := &Node{Key: 50}
	fmt.Println("Tree::", tr)
	tr.Insert(10)
	tr.Insert(200)
	tr.Insert(300)
	tr.Insert(400)
	tr.Insert(70)
	fmt.Println("Tree::", tr)
	fmt.Println("Search for 70:", tr.Search(70))
}
