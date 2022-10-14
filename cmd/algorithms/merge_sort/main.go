// Pros
// Fast. Merge sort is much faster than bubble sort, being O(n*log(n)) instead of O(n^2).
// Stable. Merge sort is also a stable sort which means that values with duplicate keys in the original list will be in the same order in the sorted list.
// Cons
// Extra memory. Most sorting algorithms can be performed using a single copy of the original array. Merge sort requires an extra array in memory to merge the sorted subarrays.
// Recursive: Merge sort requires many recursive function calls, and function calls can have significant resource overhead.

// Understand how recursions works for example if we are to get 4! it is 4*3*2*1 but how does our machine understand recursion
// It knows 4! = 3 * 2!, 2!= 2*1!, and 1!= 1. Once it figures our 1! is 1, it can go back the stack(pending multiplication) to figure out the rest.

package main

import (
	"fmt"
	"log"
)

// Runs MergeSort algorithm on a slice single
func MergeSort(slice []int) []int {

	if len(slice) < 2 { // Condition that tells recursion to stop
		return slice
	}
	mid := (len(slice)) / 2
	l := slice[:mid]
	r := slice[mid:]
	log.Printf("left: %d | right : %d \n", l, r)
	return Merge(MergeSort(l), MergeSort(r))
}

// Merges left and right slice into newly created slice
func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	count := 0

	log.Printf("size: %d | i : %d | j : %d | slice : %d | left : %d | right : %d\n", size, i, j, slice, left, right)

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			slice[count] = left[i]
			count, i = count+1, i+1
		} else {
			slice[count] = right[j]
			count, j = count+1, j+1
		}
	}
	for i < len(left) {
		slice[count] = left[i]
		count, i = count+1, i+1
	}
	for j < len(right) {
		slice[count] = right[j]
		count, j = count+1, j+1
	}
	log.Printf("Slice data %d \n", slice)
	return slice
}

func main() {
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
	s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Println("\n --- unsorted --- \n\n", s)
	fmt.Printf("\n--- sorted --- %d \n\n", MergeSort(s))
}
