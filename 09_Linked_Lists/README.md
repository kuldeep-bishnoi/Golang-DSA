# Linked Lists in Golang

## Introduction to Linked Lists
A linked list is a linear collection of data elements, whose order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a group of nodes which together represent a sequence. In its most basic form, each node contains a data field and a reference (in other words, a "link") to the next node in the sequence. This structure allows for efficient insertion or removal of elements from any position in the sequence.

In Golang, a linked list can be implemented using a struct to represent each node, with a field for the data and a pointer to the next node.

**Example: Basic Node Structure in Golang**
```go
type Node struct {
    Data int
    Next *Node
}
```

## Types of Linked Lists
There are several types of linked lists, each with its own characteristics and use cases:

### 1. Singly Linked List
In a singly linked list, each node only points to the next node in the sequence. This is the most basic type of linked list.

**Example: Singly Linked List in Golang**
```go
type SinglyLinkedList struct {
    Head *Node
}
```

### 2. Doubly Linked List
In a doubly linked list, each node not only points to the next node but also to the previous node. This allows for efficient insertion and deletion of nodes at any position in the list.

**Example: Doubly Linked List in Golang**
```go
type DoublyLinkedList struct {
    Head *Node
    Tail *Node
}

type Node struct {
    Data int
    Next *Node
    Prev *Node
}
```

### 3. Circular Linked List
In a circular linked list, the last node points back to the first node, forming a circle. This type of linked list is useful in applications where the last element needs to point back to the first element.

**Example: Circular Linked List in Golang**
```go
type CircularLinkedList struct {
    Head *Node
}

type Node struct {
    Data int
    Next *Node
}
```

## Fast and Slow Pointers
The fast and slow pointer technique is used to detect cycles in a linked list. The basic idea is to have two pointers, one moving twice as fast as the other. If there is a cycle, these two pointers will eventually meet at some node. If there is no cycle, the fast pointer will reach the end of the linked list.

**Example: Cycle Detection in Golang**
```go
func hasCycle(head *Node) bool {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

## Reversal and Recursion
Reversing a linked list can be done iteratively or recursively. The iterative approach involves initializing three pointers: previous, current, and next. The recursive approach involves calling a function recursively until the end of the list is reached, then reversing the links on the way back.

**Example: Iterative Reversal in Golang**
```go
func reverseList(head *Node) *Node {
    prev := nil
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
```

**Example: Recursive Reversal in Golang**
```go
func reverseListRecursive(head *Node) *Node {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseListRecursive(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}
```

## Doubly and Circular Linked Lists
Doubly linked lists are useful when frequent insertion and deletion of nodes are required. Circular linked lists are useful in applications where the last element needs to point back to the first element.

**Example: Insertion in Doubly Linked List in Golang**
```go
func (dll *DoublyLinkedList) InsertAtBeginning(data int) {
    newNode := &Node{Data: data}
    if dll.Head == nil {
        dll.Head = newNode
        dll.Tail = newNode
    } else {
        newNode.Next = dll.Head
        dll.Head.Prev = newNode
        dll.Head = newNode
    }
}
```

**Example: Insertion in Circular Linked List in Golang**
```go
func (cll *CircularLinkedList) InsertAtBeginning(data int) {
    newNode := &Node{Data: data}
    if cll.Head == nil {
        cll.Head = newNode
        newNode.Next = newNode
    } else {
        newNode.Next = cll.Head
        cll.Head = newNode
        cll.Head.Prev = newNode
    }
}