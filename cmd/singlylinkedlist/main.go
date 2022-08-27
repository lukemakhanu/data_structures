package main

import (
	"fmt"
	"log"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) PrintData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	log.Printf("\n")
}

func (l *linkedList) delete(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 33}
	node2 := &node{data: 14}
	node3 := &node{data: 20}
	node4 := &node{data: 49}
	node5 := &node{data: 309}
	node6 := &node{data: 12}
	node7 := &node{data: 44}
	node8 := &node{data: 59}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)
	mylist.prepend(node8)

	mylist.delete(100)
	mylist.delete(59)
	emptyList := linkedList{}
	emptyList.delete(10)
	mylist.PrintData()
}
