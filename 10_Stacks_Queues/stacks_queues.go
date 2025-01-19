package main

import "fmt"

// Stack struct using a slice
type Stack struct {
	items []int
}

// Push adds an item to the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the top item from the stack without removing it
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.items[len(s.items)-1]
}

// Queue struct using a slice
type Queue struct {
	items []int
}

// Enqueue adds an item to the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front item from the queue
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Front returns the front item from the queue without removing it
func (q *Queue) Front() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.items[0]
}

func main() {
	// Stack operations
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println("Stack Peek:", stack.Peek())
	fmt.Println("Stack Pop:", stack.Pop())
	fmt.Println("Stack Pop:", stack.Pop())
	fmt.Println("Stack Pop:", stack.Pop())

	// Queue operations
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("Queue Front:", queue.Front())
	fmt.Println("Queue Dequeue:", queue.Dequeue())
	fmt.Println("Queue Dequeue:", queue.Dequeue())
	fmt.Println("Queue Dequeue:", queue.Dequeue())
}
