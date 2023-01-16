package main

import "fmt"

type Queue struct {
	item []int
}

func (queue *Queue) Enqueue(item int) {
	queue.item = append(queue.item, item)
}
func (queue *Queue) Dequeue() int {
	data := queue.item[0]

	if len(queue.item) == 0 {
		return 0
	} else if len(queue.item) == 1 {
		queue.item = []int{}
		return data
	}
	queue.item = queue.item[1:]
	return data
}
func (queue *Queue) isEmpty() bool {
	return len(queue.item) == 0
}
func main() {
	var queue Queue
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)
	fmt.Println(queue.item)
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println(queue.item)
	fmt.Println(queue.isEmpty())
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println(queue.isEmpty())

}
