# Trees in Golang

## 📖 Simple Explanation

Think of a tree as a family tree or an organization chart. Each person (or "node") can have multiple children below them, but only one parent above them (except for the top person, who has no parent).

In computer science, trees are hierarchical data structures with a root value and subtrees of children, represented as a set of linked nodes.

```
        Root
        / | \
       /  |  \
      A   B   C
     / \     / \
    D   E   F   G
```

Trees are used to represent hierarchical relationships, organize data for quick search/insertion/deletion, and model real-world structures.

## 🔑 Key Concepts

- **Node**: Basic unit containing data and references to child nodes
- **Root**: The topmost node with no parent
- **Edge**: Connection between nodes (parent-child relationship)
- **Leaf**: Node with no children
- **Height**: Longest path from root to the furthest leaf
- **Depth**: Length of path from root to the node
- **Binary Tree**: Tree where each node has at most two children
- **Binary Search Tree (BST)**: Binary tree where left children are smaller and right children are larger than parent

## 💻 Basic Implementation in Go

### Binary Tree

```go
package main

import "fmt"

// TreeNode represents a node in a binary tree
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

// NewNode creates a new tree node with the given value
func NewNode(value int) *TreeNode {
    return &TreeNode{Value: value, Left: nil, Right: nil}
}

func main() {
    // Create a simple binary tree
    //        1
    //       / \
    //      2   3
    //     / \
    //    4   5
    
    root := NewNode(1)
    root.Left = NewNode(2)
    root.Right = NewNode(3)
    root.Left.Left = NewNode(4)
    root.Left.Right = NewNode(5)
    
    fmt.Println("In-order traversal:")
    inOrderTraversal(root)
}

// In-order traversal: left -> root -> right
func inOrderTraversal(node *TreeNode) {
    if node == nil {
        return
    }
    
    inOrderTraversal(node.Left)
    fmt.Printf("%d ", node.Value)
    inOrderTraversal(node.Right)
}
```

### Binary Search Tree

```go
// Insert a value into a BST
func (root *TreeNode) Insert(value int) *TreeNode {
    if root == nil {
        return NewNode(value)
    }
    
    if value < root.Value {
        root.Left = root.Left.Insert(value)
    } else if value > root.Value {
        root.Right = root.Right.Insert(value)
    }
    
    return root
}

// Search for a value in a BST
func (root *TreeNode) Search(value int) *TreeNode {
    if root == nil || root.Value == value {
        return root
    }
    
    if value < root.Value {
        return root.Left.Search(value)
    }
    
    return root.Right.Search(value)
}
```

## ⏱️ Time and Space Complexity

| Operation | Binary Tree | BST (Average) | BST (Worst) | AVL Tree | Notes |
|-----------|-------------|--------------|------------|----------|-------|
| Search | O(n) | O(log n) | O(n) | O(log n) | BST worst case is skewed tree |
| Insert | O(1)* | O(log n) | O(n) | O(log n) | *If we know position, otherwise O(n) |
| Delete | O(n) | O(log n) | O(n) | O(log n) | Requires finding the node first |
| Traversal | O(n) | O(n) | O(n) | O(n) | Have to visit all nodes |
| Height | O(n) | O(log n) | O(n) | O(log n) | BST worst case is linked list |
| Space | O(n) | O(n) | O(n) | O(n) | For storing the tree |

## 🧩 Common Patterns and Techniques

### Tree Traversal Methods
- **Pre-order**: Root, Left, Right
- **In-order**: Left, Root, Right (gives sorted order in BST)
- **Post-order**: Left, Right, Root
- **Level-order**: Breadth-first traversal, level by level

```go
// Pre-order traversal: root -> left -> right
func preOrderTraversal(node *TreeNode) {
    if node == nil {
        return
    }
    
    fmt.Printf("%d ", node.Value)
    preOrderTraversal(node.Left)
    preOrderTraversal(node.Right)
}

// Level-order traversal (BFS)
func levelOrderTraversal(root *TreeNode) {
    if root == nil {
        return
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        // Dequeue
        node := queue[0]
        queue = queue[1:]
        
        fmt.Printf("%d ", node.Value)
        
        // Enqueue children
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}
```

### Recursive Tree Operations
- Most tree operations can be elegantly expressed recursively
- Think of solving the problem for the current node, and let recursion handle the subtrees

```go
// Calculate height of a binary tree recursively
func height(node *TreeNode) int {
    if node == nil {
        return 0
    }
    
    leftHeight := height(node.Left)
    rightHeight := height(node.Right)
    
    // Height is the max of left and right subtree heights, plus 1 for current node
    if leftHeight > rightHeight {
        return leftHeight + 1
    }
    return rightHeight + 1
}
```

## 🔍 Real-World Applications

1. **File Systems**: Directories and files are organized in a tree structure
2. **HTML/XML DOM**: Document Object Model is a tree of elements
3. **Decision Trees**: Used in machine learning for classification and regression
4. **Network Routing**: Hierarchical routing uses tree structures
5. **Organization Charts**: Company structure representation
6. **Binary Space Partitioning**: Used in 3D computer graphics
7. **Syntax Trees**: Used in compilers to parse programming languages

## 🎯 Common Interview Questions

### Question 1: Maximum Depth of Binary Tree

**Problem**: Find the maximum depth (height) of a binary tree.

**Solution Approach**:
1. Use recursion to solve this problem
2. For each node, compute the height as 1 + max(height of left subtree, height of right subtree)
3. Base case: empty tree has height 0

```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    
    if leftDepth > rightDepth {
        return leftDepth + 1
    }
    return rightDepth + 1
}
```

**Time Complexity**: O(n) where n is the number of nodes
**Space Complexity**: O(h) where h is the height of the tree (for the recursion stack)

### Question 2: Validate Binary Search Tree

**Problem**: Determine if a binary tree is a valid binary search tree (BST).

**Solution Approach**:
1. Use recursion with range checking
2. Every node value must be within a valid range
3. For left subtree, upper bound is parent's value
4. For right subtree, lower bound is parent's value

```go
func isValidBST(root *TreeNode) bool {
    return validateBST(root, nil, nil)
}

func validateBST(node *TreeNode, min, max *int) bool {
    if node == nil {
        return true
    }
    
    // Check current node's value against bounds
    if (min != nil && node.Value <= *min) || (max != nil && node.Value >= *max) {
        return false
    }
    
    // Recursively check left and right subtrees with updated bounds
    return validateBST(node.Left, min, &node.Value) && validateBST(node.Right, &node.Value, max)
}
```

**Time Complexity**: O(n)
**Space Complexity**: O(h) for the recursion stack

## 💪 Practice Problems

1. **Easy**: [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) - Check if a binary tree is a mirror of itself
2. **Medium**: [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) - Return level order traversal
3. **Medium**: [Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/) - Find the lowest common ancestor
4. **Hard**: [Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) - Convert tree to string and back
5. **Hard**: [Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) - Find path with maximum sum

## 📚 Additional Resources

- [Visualizing Binary Trees](https://visualgo.net/en/bst) - Interactive visualization
- [Tree Traversal Explanation](https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/) - Detailed explanation of traversals
- [Binary Search Tree Tutorial](https://www.programiz.com/dsa/binary-search-tree) - Comprehensive BST explanation
- [AVL Trees](https://www.cs.usfca.edu/~galles/visualization/AVLtree.html) - AVL tree visualization
- [Red-Black Trees](https://www.cs.usfca.edu/~galles/visualization/RedBlack.html) - Red-black tree visualization

## 📝 Exercises

1. Implement a function to count the number of leaves in a binary tree
2. Write a function to find the diameter of a binary tree (longest path between any two nodes)
3. Implement a function to check if two binary trees are identical
4. Create a function to convert a sorted array to a balanced binary search tree
5. Implement a delete operation for a binary search tree