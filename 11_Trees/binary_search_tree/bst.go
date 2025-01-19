package main

import "fmt"

// Node struct for binary search tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinarySearchTree struct
type BinarySearchTree struct {
	Root *Node
}

// Insert inserts a new node into the binary search tree
func (bst *BinarySearchTree) Insert(value int) {
	bst.Root = insertNode(bst.Root, value)
}

func insertNode(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else {
		node.Right = insertNode(node.Right, value)
	}
	return node
}

// Search searches for a value in the binary search tree
func (bst *BinarySearchTree) Search(value int) bool {
	return searchNode(bst.Root, value)
}

func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	} else if value < node.Value {
		return searchNode(node.Left, value)
	} else {
		return searchNode(node.Right, value)
	}
}

func main() {
	bst := &BinarySearchTree{}
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	fmt.Println("Search 4 in BST:", bst.Search(4))
	fmt.Println("Search 6 in BST:", bst.Search(6))
}
