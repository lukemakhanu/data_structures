package main

import "fmt"

// MaxHeap : struct has a slice that holds an array
type MaxHeap struct {
	array []int
}

// Insert adds data to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array) - 1)
}

// Extract : returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.MaxHeapifyDown(0)

	return extracted
}

// MaxHeapifyDown will heapify top to bottom
func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := Left(index), Right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex { // when the left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[l] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.Swap(index, childToCompare)
			index = childToCompare
			l, r = Left(index), Right(index)
		} else {
			return
		}

	}
}

// MaxHeapifyUp will heapify from bottom up
func (h *MaxHeap) MaxHeapifyUp(index int) {
	for h.array[Parent(index)] < h.array[index] {
		h.Swap(Parent(index), index)
		index = Parent(index)
	}
}

// Parent : get the parent index
func Parent(i int) int {
	return (i - 1) / 2
}

// Left : get the left child index
func Left(i int) int {
	return 2*i + 1
}

// Right : get the right child index
func Right(i int) int {
	return 2*i + 2
}

// Swap : used to swap nodes
func (h *MaxHeap) Swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
