package main

import "fmt"

// Graph struct using adjacency matrix
type Graph struct {
	Matrix [][]int
	Size   int
}

// NewGraph creates a new graph with a given size
func NewGraph(size int) *Graph {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return &Graph{Matrix: matrix, Size: size}
}

// AddEdge adds a new edge to the graph
func (g *Graph) AddEdge(node1, node2 int) {
	g.Matrix[node1][node2] = 1
	g.Matrix[node2][node1] = 1
}

// Display prints the graph as an adjacency matrix
func (g *Graph) Display() {
	for _, row := range g.Matrix {
		fmt.Println(row)
	}
}

func main() {
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	fmt.Println("Graph Adjacency Matrix:")
	g.Display()
}
