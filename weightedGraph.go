package main

import (
	"fmt"
)

type Graph struct {
	adjMatrix [][]int
	size      int
	vertices  []string
}

func GraphInit(size int) *Graph {
	adjMatrix := make([][]int, size)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, size)
	}
	vertices := make([]string, size)
	return &Graph{size: size, adjMatrix: adjMatrix, vertices: vertices}
}

// AddVertex assigns a label to a vertex at the given index.
func (graph *Graph) AddVertex(index int, data string) {
	if index >= 0 && index < graph.size {
		graph.vertices[index] = data
	}
}

// AddEdge adds an undirected edge between two vertices with a given weight.
// It updates the adjacency matrix to reflect the connection.
func (graph *Graph) AddEdge(from int, to int, weight int) {
	if from >= 0 && to >= 0 && from <= graph.size && to <= graph.size {
		graph.adjMatrix[from][to] = weight
		graph.adjMatrix[to][from] = weight
	}
}

// Print displays the adjacency list of the graph, showing each vertex
// and its connected vertices along with the edge weights.
func (graph *Graph) Print() {
	if len(graph.vertices) == 0 {
		fmt.Println("Graph is empty.")
	}

	// Print vertices
	for from, vertex := range graph.vertices {
		fmt.Printf("%s connected with: ", vertex)
		for to, weight := range graph.adjMatrix[from] {
			if weight != 0 {
				fmt.Printf("%s weight: %d, ", graph.vertices[to], weight)
			}
		}
		fmt.Println()
	}
}
