# Stacks and Queues in Golang

## Introduction to Stacks
A stack is a linear data structure that follows the LIFO (Last In First Out) principle. This means the last element added to the stack will be the first one to be removed. Stacks are commonly used in parsing, evaluating postfix expressions, and implementing recursive algorithms.

### Basic Concepts
- **Push**: Adding an element to the top of the stack.
- **Pop**: Removing an element from the top of the stack.
- **Peek**: Looking at the top element without removing it.
- **IsEmpty**: Checking if the stack is empty.

### Implementation in Golang
Here's a basic implementation of a stack in Golang using a slice:
```go
type Stack []int

func (s *Stack) Push(value int) {
    *s = append(*s, value)
}

func (s *Stack) Pop() int {
    if s.IsEmpty() {
        return 0
    }
    value := (*s)[len(*s)-1:]
    *s = (*s)[:len(*s)-1:]
    return value[0]
}

func (s *Stack) Peek() int {
    if s.IsEmpty() {
        return 0
    }
    return (*s)[len(*s)-1:]
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}
// Example usage
stack := Stack{}
stack.Push(1)
stack.Push(2)
fmt.Println(stack.Pop()) // Output: 2
fmt.Println(stack.Peek()) // Output: 1
```
## Applications of Stacks
- **Evaluating postfix expressions**: Stacks can be used to evaluate postfix expressions by parsing the expression from left to right and using the stack to store intermediate results.
- **Implementing recursive algorithms**: Stacks can be used to implement recursive algorithms iteratively, which can be more efficient in terms of memory usage.
- **Parsing**: Stacks are used in parsing algorithms to match opening and closing brackets, parentheses, etc.

## Introduction to Queues
A queue is a linear data structure that follows the FIFO (First In First Out) principle. This means the first element added to the queue will be the first one to be removed. Queues are commonly used in job scheduling, network protocols, and handling print jobs.

### Basic Concepts
- **Enqueue**: Adding an element to the end of the queue.
- **Dequeue**: Removing an element from the front of the queue.
- **Peek**: Looking at the front element without removing it.
- **IsEmpty**: Checking if the queue is empty.

### Implementation in Golang
Here's a basic implementation of a queue in Golang using a slice:
```go
type Queue []int

func (q *Queue) Enqueue(value int) {
    *q = append(*q, value)
}

func (q *Queue) Dequeue() int {
    if q.IsEmpty() {
        return 0
    }
    value := (*q)[0]
    *q = (*q)[1:]
    return value
}

func (q *Queue) Peek() int {
    if q.IsEmpty() {
        return 0
    }
    return (*q)[0]
}

func (q *Queue) IsEmpty() bool {
    return len(*q) == 0
}
// Example usage
queue := Queue{}
queue.Enqueue(1)
queue.Enqueue(2)
fmt.Println(queue.Dequeue()) // Output: 1
fmt.Println(queue.Peek()) // Output: 2
```
## Applications of Queues
- **Job scheduling**: Queues are used to schedule jobs or tasks in the order they are received.
- **Network protocols**: Queues are used to handle network requests and responses in the order they are received.
- **Handling print jobs**: Queues are used to manage print jobs in the order they are received.

## Deque and Priority Queues
A deque (double-ended queue) is a data structure that allows adding or removing elements from both the front and the back. A priority queue is a special type of queue where elements are ordered based on their priority.

### Implementation in Golang
Here's a basic implementation of a deque in Golang using a slice:
```go
type Deque []int

func (d *Deque) AddFront(value int) {
    *d = append([]int{value}, *d...)
}

func (d *Deque) AddBack(value int) {
    *d = append(*d, value)
}

func (d *Deque) RemoveFront() int {
    if d.IsEmpty() {
        return 0
    }
    value := (*d)[0]
    *d = (*d)[1:]
    return value
}

func (d *Deque) RemoveBack() int {
    if d.IsEmpty() {
        return 0
    }
    value := (*d)[len(*d)-1:]
    *d = (*d)[:len(*d)-1:]
    return value[0]
}

func (d *Deque) IsEmpty() bool {
    return len(*d) == 0
}
// Example usage
deque := Deque{}
deque.AddFront(1)
deque.AddBack(2)
fmt.Println(deque.RemoveFront()) // Output: 1
fmt.Println(deque.RemoveBack()) // Output: 2
```

Here's a complete implementation of a priority queue in Golang using a slice:
```go
type PriorityQueue []int

func (pq *PriorityQueue) Enqueue(value int, priority int) {
    *pq = append(*pq, value)
    (*pq).heapifyUp(len(*pq)-1, priority)
}

func (pq *PriorityQueue) Dequeue() int {
    if pq.IsEmpty() {
        return 0
    }
    value := (*pq)[0]
    (*pq)[0] = (*pq)[len(*pq)-1:]
    *pq = (*pq)[:len(*pq)-1:]
    (*pq).heapifyDown(0)
    return value
}

func (pq *PriorityQueue) IsEmpty() bool {
    return len(*pq) == 0
}

func (pq *PriorityQueue) heapifyUp(index int, priority int) {
    for index > 0 {
        parentIndex := (index - 1) / 2
        if (*pq)[parentIndex] <= (*pq)[index] {
            break
        }
        (*pq)[parentIndex], (*pq)[index] = (*pq)[index], (*pq)[parentIndex]
        index = parentIndex
    }
}

func (pq *PriorityQueue) heapifyDown(index int) {
    for {
        leftChildIndex := 2*index + 1
        rightChildIndex := 2*index + 2
        largest := index

        if leftChildIndex < len(*pq) && (*pq)[leftChildIndex] > (*pq)[largest] {
            largest = leftChildIndex
        }
        if rightChildIndex < len(*pq) && (*pq)[rightChildIndex] > (*pq)[largest] {
            largest = rightChildIndex
        }
        if largest == index {
            break
        }
        (*pq)[largest], (*pq)[index] = (*pq)[index], (*pq)[largest]
        index = largest
    }
}
// Example usage
pq := PriorityQueue{}
pq.Enqueue(1, 1)
pq.Enqueue(2, 2)
pq.Enqueue(3, 3)
fmt.Println(pq.Dequeue()) // Output: 3
fmt.Println(pq.Dequeue()) // Output: 2
fmt.Println(pq.Dequeue()) // Output: 1
fmt.Println(pq.IsEmpty()) // Output: true
```
