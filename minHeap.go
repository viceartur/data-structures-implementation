package main

import "fmt"

// Min Heap
// Binary Tree implementation using Array
type HeapEntry[T any] struct {
	priority int
	message  T
}

type Heap[T any] struct {
	size     int
	elements [10]*HeapEntry[T]
}

func (heap *Heap[T]) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (heap *Heap[T]) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (heap *Heap[T]) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (heap *Heap[T]) hasLeftChild(index int) bool {
	return heap.getLeftChildIndex(index) < heap.size
}

func (heap *Heap[T]) hasRightChild(index int) bool {
	return heap.getRightChildIndex(index) < heap.size
}

func (heap *Heap[T]) hasParent(index int) bool {
	return heap.getParentIndex(index) >= 0
}

func (heap *Heap[T]) leftChild(index int) *HeapEntry[T] {
	return heap.elements[heap.getLeftChildIndex(index)]
}

func (heap *Heap[T]) rightChild(index int) *HeapEntry[T] {
	return heap.elements[heap.getRightChildIndex(index)]
}

func (heap *Heap[T]) parent(index int) *HeapEntry[T] {
	return heap.elements[heap.getParentIndex(index)]
}

func (heap *Heap[T]) swap(indexOne int, indexTwo int) {
	temp := heap.elements[indexOne]
	heap.elements[indexOne] = heap.elements[indexTwo]
	heap.elements[indexTwo] = temp
}

func (heap *Heap[T]) enlarge() {
	if heap.size == len(heap.elements) {
		newSize := len(heap.elements) * 2
		newSlice := make([]*HeapEntry[T], newSize)
		copy(newSlice, heap.elements[:])
	}
}

func (heap *Heap[T]) peek() *HeapEntry[T] {
	return heap.elements[0]
}

// poll removes the first element,
// insert the last element in the beginning, and heapify down
func (heap *Heap[T]) pop() *HeapEntry[T] {
	element := heap.peek()
	heap.elements[0] = heap.elements[heap.size-1]

	heap.size--
	heap.heapifyDown()

	return element
}

func (heap *Heap[T]) push(priority int, message T) {
	heap.enlarge()
	heap.elements[heap.size] = &HeapEntry[T]{priority: priority, message: message}
	heap.size++
	heap.heapifyUp()
}

// heapifyUp compares current element with its parrent,
// swap if parrent greater, and repeat while element has a parrent
func (heap *Heap[T]) heapifyUp() {
	index := heap.size - 1
	for heap.hasParent(index) &&
		(heap.parent(index).priority > heap.elements[index].priority) {
		heap.swap(index, heap.getParentIndex(index))
		index = heap.getParentIndex(index)
	}
}

// heapifyDown compares the current elemnt with its children (left and right),
// finds the smallest child, swaps with the child if the current is greater,
// checks for the next children
func (heap *Heap[T]) heapifyDown() {
	index := 0
	for heap.hasLeftChild(index) {
		// left child is the smallest
		smallerChildIndex := heap.getLeftChildIndex(index)

		// check whether the right child exists and it is smaller than the left child
		// if so - right child is the smallest
		if heap.hasRightChild(index) &&
			heap.rightChild(index).priority < heap.leftChild(index).priority {
			smallerChildIndex = heap.getRightChildIndex(index)
		}

		// compare the current element and the smallest
		if heap.elements[index].priority < heap.elements[smallerChildIndex].priority {
			break // we are good no swap needed - finish
		} else {
			heap.swap(index, smallerChildIndex) // swap the current element if bigger than the smallest
		}
		index = smallerChildIndex // move forward to the next element
	}
}

func (heap *Heap[T]) print() {
	for i := 0; i < heap.size; i++ {
		heapEntry := *heap.elements[i]
		fmt.Printf("Element #%d: Priority: %d, message: %v \n", i, heapEntry.priority, heapEntry.message)
	}
}
