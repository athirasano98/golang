package main

import "fmt"

type List struct {
	head *Node
}
type Node struct {
	prev *Node
	item interface{}
	next *Node
}

func (list *List) Insert(item interface{}) {
	data := &Node{nil, item, nil}
	if list.head == nil {
		list.head = data
	} else {
		p := list.head
		for p.next != nil {
			p = p.next
		}
		p.next = data
		data.prev = p

	}

}
func (list *List) Show() {
	p := list.head
	for p != nil {
		fmt.Println(p.item)
		p = p.next
	}
}

func (list *List) Delete(item interface{}) {
	p := list.head
	if p.item == item {
		list.head = p.next
		return
	}
	for p != nil {
		if p.next.item == item {
			p.next = p.next.next
			p.next.prev = p
			return
		}
	}

}

func main() {
	var list List
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)
	list.Show()
	list.Delete(10)
	fmt.Println("After deletion")
	list.Show()

}
