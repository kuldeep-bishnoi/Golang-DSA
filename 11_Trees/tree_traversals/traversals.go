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

// PreOrderTraversal performs pre-order traversal of the binary tree
func (bt *BinaryTree) PreOrderTraversal() {
	preOrder(bt.Root)
	fmt.Println()
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Print(node.Value, " ")
		preOrder(node.Left)
		preOrder(node.Right)
	}
}

// PostOrderTraversal performs post-order traversal of the binary tree
func (bt *BinaryTree) PostOrderTraversal() {
	postOrder(bt.Root)
	fmt.Println()
}

func postOrder(node *Node) {
	if node != nil {
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Print(node.Value, " ")
	}
}

func main() {
	bt := &BinaryTree{}
	bt.Root = &Node{Value: 5, Left: &Node{Value: 3, Left: &Node{Value: 2}, Right: &Node{Value: 4}}, Right: &Node{Value: 7}}
	fmt.Print("In-Order Traversal: ")
	bt.InOrderTraversal()
	fmt.Print("Pre-Order Traversal: ")
	bt.PreOrderTraversal()
	fmt.Print("Post-Order Traversal: ")
	bt.PostOrderTraversal()
}
