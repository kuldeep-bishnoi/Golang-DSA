package main

import "fmt"

// Graph struct using adjacency list
type Graph struct {
	Nodes map[int][]int
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(node int) {
	if g.Nodes == nil {
		g.Nodes = make(map[int][]int)
	}
	g.Nodes[node] = []int{}
}

// AddEdge adds a new edge to the graph
func (g *Graph) AddEdge(node1, node2 int) {
	g.Nodes[node1] = append(g.Nodes[node1], node2)
	g.Nodes[node2] = append(g.Nodes[node2], node1)
}

// DFS performs depth-first search on the graph
func (g *Graph) DFS(start int, visited map[int]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Println(start)
	for _, neighbor := range g.Nodes[start] {
		g.DFS(neighbor, visited)
	}
}

func main() {
	g := &Graph{}
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	fmt.Println("DFS Traversal:")
	visited := make(map[int]bool)
	g.DFS(1, visited)
}
