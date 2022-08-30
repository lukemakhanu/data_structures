package main

import (
	"fmt"
)

// ArraySize is the size of the array
const ArraySize = 7

// HashTable : structure
type HashTable struct {
	array [ArraySize]*bucket
}

// Bucket structure
type bucket struct {
	head *bucketNode
}

// BucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert : takes the key and adds it into a hashtable
func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.array[index].Insert(key)
}

// Search : search for key and returns bool
func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.array[index].Search(key)
}

// Delete : searches for key and deletes, returns bool
func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.array[index].Delete(key)
}

// insert will take in a key,create a node with the key and insert the node in the bucket
func (b *bucket) Insert(k string) {
	if !b.Search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Key already exists")
	}

}

// search will take in the key and return true if the key exists
func (b *bucket) Search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete : will take in the key and delete it from the bucket
func (b *bucket) Delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.key == k {
			// delete
			previousNode = previousNode.next.next
		}
		previousNode = previousNode.next
	}

}

// Hash : return hashcode
func Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashTable := Init()
	fmt.Println(testHashTable)
	fmt.Println(Hash("RANDY"))

	testBucket := &bucket{}
	testBucket.Insert("RANDY")
	testBucket.Insert("RANDY")
	testBucket.Delete("RANDY")
	fmt.Println(testBucket.Search("RANDY"))
	fmt.Println(testBucket.Search("ERIC"))

	hashtable := Init()
	list := []string{
		"ERIC",
		"KELLY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashtable.Insert(v)
	}

	hashtable.Delete("STAN")
	fmt.Println("STAN", hashtable.Search("STAN"))
	fmt.Println("KELLY", hashtable.Search("KELLY"))
}
