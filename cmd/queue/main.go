package main

import "fmt"

// Queue : holds data in the queue
type Queue struct {
	Items []int
}

// Enqueue : used to save data in the queue
func (q *Queue) Enqueue(item int) {
	q.Items = append(q.Items, item)
}

// Dequeue : removes the first item from the queue
func (q *Queue) Dequeue() int {
	if len(q.Items) > 0 {
		toRemove := q.Items[0]
		q.Items = q.Items[1:]
		return toRemove
	}
	return -1
}

func main() {
	qq := Queue{}

	qq.Enqueue(100)
	qq.Enqueue(200)
	qq.Enqueue(300)
	qq.Enqueue(400)
	qq.Enqueue(500)

	fmt.Println("Queued data | ", qq)
	dataremoved := qq.Dequeue()
	fmt.Println("removed | ", dataremoved)
	fmt.Println("Queued data after dequeue | ", qq)
}
