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

// Display prints the graph as an adjacency list
func (g *Graph) Display() {
	for node, edges := range g.Nodes {
		fmt.Printf("%d -> %v\n", node, edges)
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
	fmt.Println("Graph Adjacency List:")
	g.Display()
}
