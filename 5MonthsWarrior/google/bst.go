package main

import "fmt"

// Node is a representation of a single node in tree. (recursive ADT)
type Node struct {
	key   int
	left  *Node
	right *Node
}

// insert is a recursive method for node insertion
func (root *Node) insert(newNode *Node) {

	//if data exists, skip
	if root.key == newNode.key {
		return
	}

	// to right-subtree
	if root.key < newNode.key {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

// inOrderPrint is an in_order traversal (recursively)
func (root *Node) inOrder(fn func(n *Node)) {

	if root.left != nil {
		root.left.inOrder(fn)
	}

	fn(root)

	if root.right != nil {
		root.right.inOrder(fn)
	}

}

// Bst is a tree
// invariant:
//
// given x root, values of all left sub-tree < x, values of all right sub-tree > x
type Bst struct {
	root   *Node
	length int
}

func (bst *Bst) insert(value int) {
	node := &Node{key: value}
	bst.length++

	if bst.root == nil {
		bst.root = node
	} else {
		bst.root.insert(node)
	}
}

func (bst *Bst) inorderPrint() {
	bst.root.inOrder(func(node *Node) {
		fmt.Println(node.key)
	})
}

func main() {

	tree := new(Bst)
	tree.insert(1)
	tree.insert(0)
	tree.insert(20)
	tree.insert(4)
	tree.insert(5)
	tree.insert(6)

	tree.inorderPrint()

}
