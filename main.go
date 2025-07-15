package main

import "log"

func main() {
	// Linked List Test
	/*
		list := List[any]{}
		list.setNode("Pete")
		list.setNode("Kris")
		list.setNode("Jamie")
		list.print()
	*/
	// Priority Queue (Min Heap) Test
	/*
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
	*/

	// BST insertion
	var root *TreeNode = binarySearchTree()
	root = root.Insert(10)
	root = root.Insert(5)
	root = root.Insert(15)
	root = root.Insert(6)

	log.Println("Search 15...")
	root.Search(15)

	log.Println("Delete 10...")
	root.Delete(10)

	log.Println("Search 10...")
	root.Search(10)

}
