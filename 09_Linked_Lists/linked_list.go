package main

import "fmt"

// Node struct for singly linked list
type Node struct {
	Data int
	Next *Node
}

// SinglyLinkedList struct
type SinglyLinkedList struct {
	Head *Node
}

// InsertAtEnd inserts a new node at the end of the singly linked list
func (sll *SinglyLinkedList) InsertAtEnd(data int) {
	newNode := &Node{Data: data}
	if sll.Head == nil {
		sll.Head = newNode
		return
	}
	current := sll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Display prints all elements in the singly linked list
func (sll *SinglyLinkedList) Display() {
	current := sll.Head
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

// Node struct for doubly linked list
type DNode struct {
	Data int
	Next *DNode
	Prev *DNode
}

// DoublyLinkedList struct
type DoublyLinkedList struct {
	Head *DNode
	Tail *DNode
}

// InsertAtEnd inserts a new node at the end of the doubly linked list
func (dll *DoublyLinkedList) InsertAtEnd(data int) {
	newNode := &DNode{Data: data}
	if dll.Head == nil {
		dll.Head = newNode
		dll.Tail = newNode
		return
	}
	dll.Tail.Next = newNode
	newNode.Prev = dll.Tail
	dll.Tail = newNode
}

// Display prints all elements in the doubly linked list
func (dll *DoublyLinkedList) Display() {
	current := dll.Head
	for current != nil {
		fmt.Print(current.Data, " <-> ")
		current = current.Next
	}
	fmt.Println("nil")
}

// Node struct for circular linked list
type CNode struct {
	Data int
	Next *CNode
}

// CircularLinkedList struct
type CircularLinkedList struct {
	Head *CNode
}

// InsertAtEnd inserts a new node at the end of the circular linked list
func (cll *CircularLinkedList) InsertAtEnd(data int) {
	newNode := &CNode{Data: data}
	if cll.Head == nil {
		cll.Head = newNode
		newNode.Next = cll.Head
		return
	}
	current := cll.Head
	for current.Next != cll.Head {
		current = current.Next
	}
	current.Next = newNode
	newNode.Next = cll.Head
}

// Display prints all elements in the circular linked list
func (cll *CircularLinkedList) Display() {
	if cll.Head == nil {
		fmt.Println("nil")
		return
	}
	current := cll.Head
	for {
		fmt.Print(current.Data, " -> ")
		current = current.Next
		if current == cll.Head {
			break
		}
	}
	fmt.Println("(head)")
}

func main() {
	// Singly Linked List
	sll := &SinglyLinkedList{}
	sll.InsertAtEnd(1)
	sll.InsertAtEnd(2)
	sll.InsertAtEnd(3)
	fmt.Print("Singly Linked List: ")
	sll.Display()

	// Doubly Linked List
	dll := &DoublyLinkedList{}
	dll.InsertAtEnd(1)
	dll.InsertAtEnd(2)
	dll.InsertAtEnd(3)
	fmt.Print("Doubly Linked List: ")
	dll.Display()

	// Circular Linked List
	cll := &CircularLinkedList{}
	cll.InsertAtEnd(1)
	cll.InsertAtEnd(2)
	cll.InsertAtEnd(3)
	fmt.Print("Circular Linked List: ")
	cll.Display()
}
