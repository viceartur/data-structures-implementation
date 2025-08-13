package main

import (
	"fmt"
	"slices"
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

// dfsSearch is a helper function for performing a depth-first search (DFS) recursively.
// It marks the current vertex as visited and prints its label, then recursively visits all
// adjacent unvisited vertices.
func (graph *Graph) dfsSearch(index int, visited []bool) {
	visited[index] = true
	fmt.Printf("%s, ", graph.vertices[index])

	for i := range graph.size {
		if graph.adjMatrix[index][i] != 0 && !visited[i] {
			graph.dfsSearch(i, visited)
		}
	}
}

// DFS starts a depth-first search (DFS) traversal from the vertex with the given label.
// It initializes the visited slice and finds the start index, then calls dfsSearch.
func (graph *Graph) DFS(startVertex string) {
	visited := make([]bool, graph.size)
	startIndex := slices.Index(graph.vertices, startVertex)
	graph.dfsSearch(startIndex, visited)
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
