package main

import (
	"fmt"
)

// MaxHeap : keep track of data
type MaxHeap struct {
	array []BetDetails
}

// BetDetails : bet details
type BetDetails struct {
	BetID     int
	Priority  int
	BetAmount float32
}

// MaxHeap1 : struct has a slice that holds an array
type MaxHeap1 struct {
	array []int
}

// Insert adds data to the heap
func (h *MaxHeap) Insert(key BetDetails) {
	//h.array = append(h.array, key)
	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array)-1, key.Priority)
}

// Extract : returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}
	h.array[0] = h.array[l] // reassign the index
	h.array = h.array[:l]   // reduce the length of the array.

	h.MaxHeapifyDown(0)

	return extracted.Priority
}

// MaxHeapifyDown will heapify top to bottom
func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := Left(index), Right(index)
	childToCompare := 0

	for l <= lastIndex {

		if l == lastIndex { // when the left child is the only child
			childToCompare = l
		} else {

			leftChild := h.array[l]
			rightChild := h.array[r]

			if leftChild.Priority > rightChild.Priority {
				// when left child is larger
				childToCompare = l
			} else {
				// when right child is larger
				childToCompare = r
			}
		}

		// compare array value of current index to larger child and swap if smaller
		xx := h.array[index]
		tt := h.array[childToCompare]

		if xx.Priority < tt.Priority {
			h.Swap(index, childToCompare)
			index = childToCompare
			l, r = Left(index), Right(index)
		} else {
			return
		}

	}
}

// MaxHeapifyUp will heapify from bottom up
func (h *MaxHeap) MaxHeapifyUp(index int, Priority int) {

	for x, dt := range h.array {
		fmt.Println("x :", x, "Priority : ", dt.Priority, "BetID : ", dt.BetID, "index :", index)

		parentIndx := Parent(index)
		dx := h.array[parentIndx]
		dd := h.array[index]
		if dx.Priority < dd.Priority {
			// Swap
			h.Swap(parentIndx, index)
			index = parentIndx
		}

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
	buildHeap := HeapData()
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	fmt.Println("start extracting")

	for i := 0; i < len(buildHeap); i++ {
		m.Extract()
		fmt.Println(m)
	}
}

// HeapData : retrurn heap data
func HeapData() []BetDetails {
	var al MaxHeap

	g := BetDetails{
		BetID:     1,
		Priority:  10,
		BetAmount: 10.00,
	}
	al.array = append(al.array, g)

	g1 := BetDetails{
		BetID:     2,
		Priority:  20,
		BetAmount: 310.00,
	}
	al.array = append(al.array, g1)

	g2 := BetDetails{
		BetID:     3,
		Priority:  30,
		BetAmount: 100.00,
	}
	al.array = append(al.array, g2)

	g3 := BetDetails{
		BetID:     4,
		Priority:  5,
		BetAmount: 370.00,
	}
	al.array = append(al.array, g3)

	g4 := BetDetails{
		BetID:     5,
		Priority:  7,
		BetAmount: 1230.00,
	}
	al.array = append(al.array, g4)

	g5 := BetDetails{
		BetID:     6,
		Priority:  9,
		BetAmount: 400.00,
	}
	al.array = append(al.array, g5)

	g6 := BetDetails{
		BetID:     7,
		Priority:  11,
		BetAmount: 890.00,
	}
	al.array = append(al.array, g6)

	g7 := BetDetails{
		BetID:     8,
		Priority:  13,
		BetAmount: 890.00,
	}
	al.array = append(al.array, g7)

	g8 := BetDetails{
		BetID:     9,
		Priority:  15,
		BetAmount: 890.00,
	}
	al.array = append(al.array, g8)

	g11 := BetDetails{
		BetID:     9,
		Priority:  17,
		BetAmount: 890.00,
	}
	al.array = append(al.array, g11)

	return al.array
}
