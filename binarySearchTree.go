package main

import "log"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func binarySearchTree() *TreeNode {
	return &TreeNode{}
}

// Insert method recursively calls itself and adds the new node where needed.
func (node *TreeNode) Insert(data int) *TreeNode {
	// Base case - node doesn't exist
	if node == nil {
		return &TreeNode{data: data}
	}

	if data < node.data {
		// Update pointer to the left node
		node.left = node.left.Insert(data)
	} else if data > node.data {
		// Update pointer to the right node
		node.right = node.right.Insert(data)
	}

	return node
}

// Search looks for a node with the given data and returns it if found, otherwise returns nil.
func (node *TreeNode) Search(data int) *TreeNode {
	if node == nil {
		log.Println("Node not found.")
		return nil
	}

	if data < node.data {
		return node.left.Search(data)
	} else if data > node.data {
		return node.right.Search(data)
	} else {
		log.Println("Node found:", data)
		return node
	}
}

// MinValueNode looks for the minival value node for the Root Node.
func (node *TreeNode) MinValueNode() *TreeNode {
	minNode := node
	for minNode != nil && minNode.left != nil {
		minNode = minNode.left
	}
	return minNode
}

// Delete a Node
// 1. If the node is a leaf node, remove it by removing the link to it.
// 2. If the node only has one child node, connect the parent node of the remove node to that child node.
// 3. If the node has both right and left child nodes: Find the node's in-order successor, change values with that node, then delete it.
func (node *TreeNode) Delete(data int) *TreeNode {
	if node == nil {
		log.Println("Node not found.")
		return nil
	}

	if data < node.data {
		node.left = node.left.Delete(data)
	} else if data > node.data {
		node.right = node.right.Delete(data)
	} else {
		// Node found

		// Node with only one child or no child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// Node with two children, get the in-order successor
		minNode := node.right.MinValueNode()
		node.data = minNode.data                     // replace delete node with successor
		node.right = node.right.Delete(minNode.data) // find and delete successor
	}

	return node
}
