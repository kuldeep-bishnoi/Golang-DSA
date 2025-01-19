# Graphs in Golang

## Introduction to Graphs
A graph is a data structure that consists of a set of nodes (vertices) and a set of edges that connect these nodes. Graphs are used to represent relationships between objects. In Golang, a graph can be implemented using an adjacency list or an adjacency matrix.

### Basic Concepts
- **Node (Vertex)**: A node is a basic unit of a graph. It contains a value and a list of its adjacent nodes.
- **Edge**: An edge is a connection between two nodes. It can be directed (one-way) or undirected (two-way).
- **Weighted Edge**: A weighted edge has a value associated with it, which can represent the cost or distance between nodes.

**Example: Basic Node Structure in Golang**
```go
type Node struct {
    Value int
    Adjacent []*Node
}
```

### Graph Representations
- **Adjacency List**: In an adjacency list, each node has a list of its adjacent nodes. This representation is efficient for sparse graphs.
- **Adjacency Matrix**: In an adjacency matrix, a two-dimensional array is used to represent the connections between nodes. This representation is efficient for dense graphs.

**Example: Adjacency List in Golang**
```go
type Graph struct {
    Nodes []*Node
}

func (g *Graph) AddNode(value int) {
    node := &Node{Value: value}
    g.Nodes = append(g.Nodes, node)
}

func (g *Graph) AddEdge(from, to int) {
    g.Nodes[from].Adjacent = append(g.Nodes[from].Adjacent, g.Nodes[to])
}
```

## Depth-First Search (DFS)
Depth-First Search is a graph traversal algorithm that starts at a node and explores as far as possible along each branch before backtracking. It uses a stack to keep track of the nodes to visit.

**Example: Depth-First Search in Golang**
```go
func DFS(node *Node, visited map[*Node]bool) {
    if node == nil {
        return
    }
    visited[node] = true
    fmt.Println(node.Value)
    for _, adjacent := range node.Adjacent {
        if !visited[adjacent] {
            DFS(adjacent, visited)
        }
    }
}
```

## Breadth-First Search (BFS)
Breadth-First Search is a graph traversal algorithm that starts at a node and explores all its adjacent nodes before moving on to the next level. It uses a queue to keep track of the nodes to visit.

**Example: Breadth-First Search in Golang**
```go
func BFS(node *Node, visited map[*Node]bool) {
    if node == nil {
        return
    }
    queue := []*Node{node}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        if !visited[current] {
            visited[current] = true
            fmt.Println(current.Value)
            for _, adjacent := range current.Adjacent {
                if !visited[adjacent] {
                    queue = append(queue, adjacent)
                }
            }
        }
    }
}
```

## Shortest Path Algorithms
- **Dijkstra's Algorithm**: Dijkstra's algorithm is a greedy algorithm that finds the shortest path between nodes in a graph. It works for both directed and undirected graphs, but it does not work for graphs with negative edge weights.
- **Bellman-Ford Algorithm**: Bellman-Ford algorithm is a dynamic programming algorithm that finds the shortest path between nodes in a graph. It can handle graphs with negative edge weights, but it is slower than Dijkstra's algorithm.

**Example: Dijkstra's Algorithm in Golang**
```go
func Dijkstra(graph *Graph, start *Node) map[*Node]int {
    distances := make(map[*Node]int)
    for _, node := range graph.Nodes {
        distances[node] = math.MaxInt32
    }
    distances[start] = 0
    queue := make(PriorityQueue, 0)
    heap.Push(&queue, &Item{value: start, priority: 0})
    for len(queue) > 0 {
        current := heap.Pop(&queue).(*Item).value
        for _, adjacent := range current.Adjacent {
            distance := distances[current] + 1
            if distance < distances[adjacent] {
                distances[adjacent] = distance
                heap.Push(&queue, &Item{value: adjacent, priority: distance})
            }
        }
    }
    return distances
}
```

## Graph Problems and Applications
Graphs have numerous applications in computer science, including:

- **Social Networks**: Graphs are used to represent social networks, with nodes representing people and edges representing relationships.
- **Network Routing**: Graphs are used to represent computer networks, with nodes representing routers and edges representing connections.
- **Recommendation Systems**: Graphs are used to recommend items to users based on their preferences and the preferences of similar users.
- **Game Development**: Graphs are used to represent game maps, with nodes representing locations and edges representing paths between them.

**Example: Finding the Shortest Path in a Graph in Golang**
```go
func findShortestPath(graph *Graph, start, end *Node) int {
    distances := Dijkstra(graph, start)
    return distances[end]
}
```