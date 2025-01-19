# Trees in Golang

## Introduction to Binary Trees
A binary tree is a data structure in which each node has at most two children (i.e., left child and right child). This structure is useful for organizing and searching data efficiently.

### Basic Concepts
- **Node**: A node is the basic unit of a binary tree. It contains a value and references to its left and right children.
- **Root Node**: The root node is the topmost node of the binary tree.
- **Leaf Node**: A leaf node is a node with no children.
- **Height**: The height of a binary tree is the number of edges on the longest path from the root to a leaf.
- **Depth**: The depth of a node is the number of edges from the root to the node.

### Types of Binary Trees
- **Full Binary Tree**: A full binary tree is a tree in which every node has exactly two children.
- **Empty Binary Tree**: An empty binary tree is a tree with no nodes.
- **Complete Binary Tree**: A complete binary tree is a tree in which every level, except possibly the last, is completely filled, and all nodes are as far left as possible.
- **Balanced Binary Tree**: A balanced binary tree is a tree in which the height of the left and right subtrees of every node differs by at most one.

**Example: Basic Node Structure in Golang**
```go
type Node struct {
    Value int
    Left  *Node
    Right *Node
}
```

## Binary Tree Traversals
Binary tree traversals are algorithms for visiting each node in a binary tree exactly once. There are three main types of traversals: in-order, pre-order, and post-order.

### In-Order Traversal
In an in-order traversal, the left subtree is visited first, then the root, and finally the right subtree.

**Example: In-Order Traversal in Golang**
```go
func inOrderTraversal(node *Node) {
    if node == nil {
        return
    }
    inOrderTraversal(node.Left)
    fmt.Println(node.Value)
    inOrderTraversal(node.Right)
}
```

### Pre-Order Traversal
In a pre-order traversal, the root is visited first, then the left subtree, and finally the right subtree.

**Example: Pre-Order Traversal in Golang**
```go
func preOrderTraversal(node *Node) {
    if node == nil {
        return
    }
    fmt.Println(node.Value)
    preOrderTraversal(node.Left)
    preOrderTraversal(node.Right)
}
```

### Post-Order Traversal
In a post-order traversal, the left and right subtrees are visited first, and then the root.

**Example: Post-Order Traversal in Golang**
```go
func postOrderTraversal(node *Node) {
    if node == nil {
        return
    }
    postOrderTraversal(node.Left)
    postOrderTraversal(node.Right)
    fmt.Println(node.Value)
}
```

## Binary Search Trees (BST)
A binary search tree is a binary tree in which each node has a comparable value; for any given node, all elements in its left subtree are less than the node, and all the elements in its right subtree are greater than the node.

### Concepts and Operations
- **Insertion**: Inserting a new node into the BST while maintaining the BST property.
- **Deletion**: Deleting a node from the BST while maintaining the BST property.
- **Search**: Finding a node with a given value in the BST.

**Example: Insertion in BST in Golang**
```go
func insertNode(root **Node, value int) {
    if *root == nil {
        *root = &Node{Value: value}
        return
    }
    if value < (*root).Value {
        insertNode(&((*root).Left), value)
    } else {
        insertNode(&((*root).Right), value)
    }
}
```

## Balanced Binary Trees
A balanced binary tree is a tree in which the height of the left and right subtrees of every node differs by at most one. This property ensures that search, insertion, and deletion operations can be performed efficiently.

### AVL and Red-Black Trees
- **AVL Tree**: An AVL tree is a self-balancing binary search tree that ensures the height of the tree remains relatively small by rotating nodes when the balance factor becomes too large.
- **Red-Black Tree**: A red-black tree is a self-balancing binary search tree with a guarantee of O(log n) time for search, insert, and delete operations.

**Example: AVL Tree Node Structure in Golang**
```go
type AVLNode struct {
    Value    int
    Height   int
    Left     *AVLNode
    Right    *AVLNode
    Balance  int
}
```

## Tree Problems and Applications
Trees have numerous applications in computer science, including:

- **File System Organization**: Trees are used to represent the hierarchical structure of file systems.
- **Database Indexing**: Trees are used in database indexing to facilitate fast lookup and retrieval of data.
- **Compilers**: Trees are used to represent the parse trees of source code in compilers.
- **Web Page Navigation**: Trees are used to represent the hierarchical structure of web pages for navigation.

**Example: Finding the Height of a Tree in Golang**
```go
func findHeight(node *Node) int {
    if node == nil {
        return 0
    }
    return 1 + max(findHeight(node.Left), findHeight(node.Right))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```