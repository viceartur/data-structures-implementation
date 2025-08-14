package main

import (
	"fmt"
)

func main() {
	// Linked List
	list := List[any]{}
	list.setNode("Pete")
	list.setNode("Kris")
	list.setNode("Jamie")
	list.print()

	// Priority Queue (Min Heap)
	heap := Heap[string]{}
	heap.push(100, "this is hundred alright?")
	heap.push(50, "should be faster than hundred")
	heap.push(200, "bigger than hundred for sure")
	heap.push(1000, "this is the last one")
	heap.push(1, "THE FIRST!!!111")
	heap.push(2, "im the second one :)")
	heap.print()
	fmt.Printf("top dequeued: %s \n", heap.pop().message)
	heap.print()

	// BST
	var root *TreeNode
	root = root.Insert(10)
	root = root.Insert(5)
	root = root.Insert(15)
	root = root.Insert(6)

	fmt.Println("In-Order traversal:")
	root.InOrderTraversal()

	fmt.Println("Pre-Order traversal:")
	root.PreOrderTraversal()

	fmt.Println("Post-Order traversal:")
	root.PostOrderTraversal()

	fmt.Println("Search 15...")
	root.Search(15)

	fmt.Println("Delete 10...")
	root.Delete(10)

	fmt.Println("Search 10...")
	root.Search(10)

	// AVL Tree
	var rootAVL *AVLTreeNode
	rootAVL = rootAVL.Insert(5)
	rootAVL = rootAVL.Insert(6)
	rootAVL = rootAVL.Insert(7)
	rootAVL = rootAVL.Insert(8)
	rootAVL = rootAVL.Insert(9)
	rootAVL = rootAVL.Insert(10)

	fmt.Println("In-Order traversal:")
	rootAVL.InOrderTraversal()
	fmt.Println("Deleting 7...")
	rootAVL.Delete(7)
	fmt.Println("In-Order traversal:")
	rootAVL.InOrderTraversal()

	// Graph
	graph := GraphInit(4)
	graph.AddVertex(0, "A")
	graph.AddVertex(1, "B")
	graph.AddVertex(2, "C")
	graph.AddVertex(3, "D")
	graph.AddEdge(0, 1, 3) // A - B weight 3
	graph.AddEdge(0, 2, 2) // A - C weight 2
	graph.AddEdge(0, 3, 4) // A - D weight 4
	graph.AddEdge(1, 2, 1) // B - C weight 1
	graph.Print()
	graph.DFS("C")
	graph.BFS("A")
}
