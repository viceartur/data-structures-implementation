package main

import "fmt"

type AVLTreeNode struct {
	data   int
	left   *AVLTreeNode
	right  *AVLTreeNode
	height int
}

func (node *AVLTreeNode) GetHeight() int {
	if node == nil {
		return 0
	}
	return node.height
}

func (node *AVLTreeNode) GetBalance() int {
	if node == nil {
		return 0
	}

	return node.left.GetHeight() - node.right.GetHeight()
}

// RightRotate performs a right rotation on the current node (y).
// This is used to fix Left-Left (LL) and Left-Right (LR) imbalances.
//
//	    y                 x
//	   / \               / \
//	  x   T3   =>       T1  y
//	 / \                   / \
//	T1 T2                 T2 T3
func (y *AVLTreeNode) RightRotate() *AVLTreeNode {
	if y == nil || y.left == nil {
		return y // nothing to rotate
	}
	x := y.left
	T2 := x.right

	// Perform rotation
	x.right = y
	y.left = T2

	// Update heights
	y.height = 1 + max(y.left.GetHeight(), y.right.GetHeight())
	x.height = 1 + max(x.left.GetHeight(), x.right.GetHeight())

	return x
}

// LeftRotate performs a left rotation on the current node (x).
// This is used to fix Right-Right (RR) and Right-Left (RL) imbalances.
//
//	  x                    y
//	 / \                  / \
//	T1  y     =>         x  T3
//	   / \              / \
//	  T2 T3            T1 T2
func (x *AVLTreeNode) LeftRotate() *AVLTreeNode {
	if x == nil || x.right == nil {
		return x // nothing to rotate
	}
	y := x.right
	T2 := y.left

	// Perform rotation
	y.left = x
	x.right = T2

	// Update heights
	x.height = 1 + max(x.left.GetHeight(), x.right.GetHeight())
	y.height = 1 + max(y.left.GetHeight(), y.right.GetHeight())

	return y
}

func (node *AVLTreeNode) Insert(data int) *AVLTreeNode {
	if node == nil {
		return &AVLTreeNode{data: data}
	}

	if data < node.data {
		node.left = node.left.Insert(data)
	} else if data > node.data {
		node.right = node.right.Insert(data)
	}

	// Update the balance factor and balance the tree
	node.height = 1 + max(node.left.GetHeight(), node.right.GetHeight())
	balance := node.GetBalance()

	// Left-Left
	if balance > 1 && node.left.GetBalance() >= 0 {
		return node.RightRotate()
	}

	// Left-Right
	if balance > 1 && node.left.GetBalance() < 0 {
		node.left.LeftRotate()
		return node.RightRotate()
	}

	// Right-Right
	if balance < -1 && node.right.GetBalance() <= 0 {
		return node.LeftRotate()
	}

	// Right-Left
	if balance < -1 && node.right.GetBalance() > 0 {
		node.right.RightRotate()
		return node.LeftRotate()
	}

	return node
}

func (node *AVLTreeNode) InOrderTraversal() {
	if node == nil {
		return
	}
	node.left.InOrderTraversal()
	fmt.Println(node.data)
	node.right.InOrderTraversal()
}

func (node *AVLTreeNode) MinValueNode() *AVLTreeNode {
	minNode := node
	for minNode != nil && minNode.left != nil {
		minNode = minNode.left
	}
	return minNode
}

func (node *AVLTreeNode) Delete(data int) *AVLTreeNode {
	if node == nil {
		fmt.Println("Node not found.")
		return nil
	}

	if data < node.data {
		node.left = node.left.Delete(data)
	} else if data > node.data {
		node.right = node.right.Delete(data)
	} else {

		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		minNode := node.right.MinValueNode()
		node.data = minNode.data
		node.right = node.right.Delete(minNode.data)
	}

	// Update the balance factor and balance the tree
	node.height = 1 + max(node.left.GetHeight(), node.right.GetHeight())
	balance := node.GetBalance()

	// Left-Left
	if balance > 1 && node.left.GetBalance() >= 0 {
		return node.RightRotate()
	}

	// Left-Right
	if balance > 1 && node.left.GetBalance() < 0 {
		node.left.LeftRotate()
		return node.RightRotate()
	}

	// Right-Right
	if balance < -1 && node.right.GetBalance() <= 0 {
		return node.LeftRotate()
	}

	// Right-Left
	if balance < -1 && node.right.GetBalance() > 0 {
		node.right.RightRotate()
		return node.LeftRotate()
	}

	return node
}
