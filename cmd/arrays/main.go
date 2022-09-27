package main

import "fmt"

func main() {
	fmt.Printf("Array datastructures")
	array := []int{1, 2, 4, 8, 7, 9}
	for i, d := range array {
		fmt.Printf("index : %d data %d \n", i, d)
	}
	// search
	k := 7
	data, status := findElement(array, k)
	if status == true {
		fmt.Printf("value found at index : %d \n", data)
	} else {
		fmt.Printf("value not found \n")
	}
	// Add at the end
	d := 12
	dd := addAtEnd(array, d)
	fmt.Printf("after add : %v \n", dd)
	// Add a position
	val := 88
	pos := 2
	tt := insertAtPos(array, val, pos)
	fmt.Printf("add at position : %v \n", tt)
}

func findElement(arr []int, key int) (int, bool) {
	for k, v := range arr {
		if v == key {
			return k, true
		}
	}
	return -1, false
}

func addAtEnd(arr []int, key int) []int {
	arr = append(arr, key)
	return arr
}

func insertAtPos(arr []int, key int, pos int) []int {
	//val := arr[pos]
	newSlice := make([]int, pos+1)
	copy(newSlice, arr[:pos])
	fmt.Println("newSlice before update:>>", newSlice)
	newSlice[pos] = key
	fmt.Println("newSlice:", newSlice)
	fmt.Println("slice:", arr)

	arr = append(newSlice, arr[pos:]...)
	fmt.Println("slice ***:", arr)

	return arr
}
