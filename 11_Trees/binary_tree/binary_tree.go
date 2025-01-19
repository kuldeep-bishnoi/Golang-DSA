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

func main() {
	bt := &BinaryTree{}
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(2)
	bt.Insert(4)
	fmt.Print("Binary Tree In-Order Traversal: ")
	bt.InOrderTraversal()
}
