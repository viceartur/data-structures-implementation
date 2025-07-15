package main

import "log"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type List[T any] struct {
	firstNode *Node[T]
	lastNode  *Node[T]
}

func (list *List[T]) setNode(val T) bool {
	node := &Node[T]{value: val}

	if list.firstNode == nil {
		list.firstNode = node
		list.lastNode = node
		return true
	}

	list.lastNode.next = node
	list.lastNode = node

	return true
}

func (list *List[T]) print() {
	currNode := list.firstNode

	for currNode != nil {
		log.Println("Element: ", currNode.value)
		currNode = currNode.next
	}
}
