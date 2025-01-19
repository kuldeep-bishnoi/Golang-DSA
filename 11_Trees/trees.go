package main

import "fmt"

// Node struct for binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinaryTree struct
type BinaryTree struct {
	Root *Node
}

// Insert inserts a new node into the binary tree
func (bt *BinaryTree) Insert(value int) {
	bt.Root = insertNode(bt.Root, value)
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

// InOrderTraversal performs in-order traversal of the binary tree
func (bt *BinaryTree) InOrderTraversal() {
	inOrder(bt.Root)
	fmt.Println()
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.Left)
		fmt.Print(node.Value, " ")
		inOrder(node.Right)
	}
}

// BinarySearchTree struct
type BinarySearchTree struct {
	Root *Node
}

// Insert inserts a new node into the binary search tree
func (bst *BinarySearchTree) Insert(value int) {
	bst.Root = insertNode(bst.Root, value)
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
	// Binary Tree
	bt := &BinaryTree{}
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(2)
	bt.Insert(4)
	fmt.Print("Binary Tree In-Order Traversal: ")
	bt.InOrderTraversal()

	// Binary Search Tree
	bst := &BinarySearchTree{}
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(2)
	bst.Insert(4)
	fmt.Print("Binary Search Tree In-Order Traversal: ")
	bt.InOrderTraversal()
	fmt.Println("Search 4 in BST:", bst.Search(4))
	fmt.Println("Search 6 in BST:", bst.Search(6))
}
