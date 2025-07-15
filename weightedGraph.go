package main

type Vertex[T any] struct {
	value T
}

type Graph[T any] struct {
	edges [][]T
	
}