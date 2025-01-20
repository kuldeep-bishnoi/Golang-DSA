# Binary Trees in Golang

## Introduction to Binary Trees
A binary tree is a hierarchical data structure in which each node has at most two children, referred to as the left child and the right child. Binary trees are used to implement binary search trees and binary heaps, and they are the basis for more advanced data structures.

### Basic Concepts
- **Node**: The basic unit of a binary tree, containing a value and references to its left and right children.
- **Root Node**: The topmost node of the tree.
- **Leaf Node**: A node with no children.
- **Height**: The number of edges on the longest path from the root to a leaf.
- **Depth**: The number of edges from the root to a node.

### Diagram of a Binary Tree
```
        1
       / \
      2   3
     / \ / \
    4  5 6  7
```

## Types of Binary Trees
- **Full Binary Tree**: Every node has either 0 or 2 children.
  - Example: A tree where each node has exactly two children or none.
- **Complete Binary Tree**: All levels are completely filled except possibly the last, which is filled from left to right.
  - Example: A tree where all levels are filled except the last, which is filled from left to right.
- **Balanced Binary Tree**: The height of the left and right subtrees of any node differ by at most one.
  - Example: AVL trees and Red-Black trees.

## Binary Tree Operations
### Insertion
Inserting a new node into a binary tree involves finding the correct position for the node while maintaining the properties of the tree.
- **Algorithm**:
  1. Start at the root.
  2. Compare the value to be inserted with the current node's value.
  3. If the value is less, move to the left child; if greater, move to the right child.
  4. Repeat until the correct position is found.

### Deletion
Deleting a node from a binary tree requires rearranging the tree to maintain its properties.
- **Algorithm**:
  1. Find the node to be deleted.
  2. If the node has no children, simply remove it.
  3. If the node has one child, replace it with its child.
  4. If the node has two children, find the in-order successor, replace the node with the successor, and delete the successor.

### Traversal
- **In-Order Traversal**: Visit the left subtree, the root, and then the right subtree.
  - Example: Produces a sorted list of values for binary search trees.
- **Pre-Order Traversal**: Visit the root, the left subtree, and then the right subtree.
  - Example: Used to create a copy of the tree.
- **Post-Order Traversal**: Visit the left subtree, the right subtree, and then the root.
  - Example: Used to delete the tree.

## Applications of Binary Trees
- **Expression Trees**: Used in compilers to represent expressions.
  - Example: Parsing mathematical expressions.
- **Huffman Coding Trees**: Used in data compression algorithms.
  - Example: Huffman coding for lossless data compression.
- **Binary Search Trees**: Used for efficient searching and sorting.
  - Example: Implementing dynamic sets and lookup tables.

Binary trees are fundamental data structures that provide the foundation for more complex structures and algorithms. Understanding their properties and operations is crucial for solving various computational problems efficiently.
